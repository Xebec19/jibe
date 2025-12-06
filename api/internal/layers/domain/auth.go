package domain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type CreateNonceDTO struct {
	Eth_Addr string `json:"eth_addr"`
}

type VerifyRequest struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type VerifyResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

// ParseSIWEMessage parses a SIWE message string
type SIWEMessage struct {
	Domain         string
	Address        string
	Statement      string
	URI            string
	Version        string
	ChainID        string
	Nonce          string
	IssuedAt       string
	ExpirationTime string
	NotBefore      string
	RequestID      string
	Resources      []string
}

func ParseSIWEMessage(message string) (*SIWEMessage, error) {
	siwe := &SIWEMessage{}

	lines := strings.Split(message, "\n")
	if len(lines) < 4 {
		return nil, fmt.Errorf("invalid SIWE message format")
	}

	// Parse domain and address
	domainRegex := regexp.MustCompile(`^(.+) wants you to sign in with your Ethereum account:`)
	matches := domainRegex.FindStringSubmatch(lines[0])
	if len(matches) > 1 {
		siwe.Domain = matches[1]
	}

	siwe.Address = strings.TrimSpace(lines[1])

	// Skip empty line and statement
	statementStart := 3
	statementEnd := statementStart
	for i := statementStart; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" || strings.HasPrefix(lines[i], "URI:") {
			statementEnd = i
			break
		}
	}

	if statementEnd > statementStart {
		siwe.Statement = strings.Join(lines[statementStart:statementEnd], "\n")
	}

	// Parse remaining fields
	for i := statementEnd; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			continue
		}

		key, value := parts[0], parts[1]
		switch key {
		case "URI":
			siwe.URI = value
		case "Version":
			siwe.Version = value
		case "Chain ID":
			siwe.ChainID = value
		case "Nonce":
			siwe.Nonce = value
		case "Issued At":
			siwe.IssuedAt = value
		case "Expiration Time":
			siwe.ExpirationTime = value
		case "Not Before":
			siwe.NotBefore = value
		case "Request ID":
			siwe.RequestID = value
		}
	}

	return siwe, nil
}

func VerifySignature(message, signature, expectedAddress string) (bool, error) {
	// Add Ethereum signed message prefix
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefix))

	// Decode signature
	sig, err := hexutil.Decode(signature)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}

	if len(sig) != 65 {
		return false, fmt.Errorf("invalid signature length")
	}

	// Adjust V value (EIP-155)
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	// Recover public key
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// Get address from public key
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// Compare addresses (case-insensitive)
	return strings.EqualFold(recoveredAddr.Hex(), expectedAddress), nil
}

// GenerateRefreshToken function to create a secure random refresh token
func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32) // 256 bits
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// HashToken function to create SHA-256 hash of a string
func HashToken(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

type DeviceInfo struct {
	UserAgent string
	IP        string
	Platform  string
}

func GetDeviceInfo(r *http.Request) DeviceInfo {
	// Get IP address
	ip := r.Header.Get("X-Forwarded-For") // if behind proxy/load balancer
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}

	return DeviceInfo{
		UserAgent: r.Header.Get("User-Agent"),
		IP:        ip,
		Platform:  r.Header.Get("Sec-Ch-Ua-Platform"), // modern browsers send this
	}
}

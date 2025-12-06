package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenJWTClaims struct {
	Iss string
	Sub string
	Aud string
	Exp time.Time
	Iat time.Time
	Nbf time.Time
	Jti string
}

// Token expects jwtclaims and secret, and create a jwt token
func Token(claims TokenJWTClaims, secret []byte) (string, error) {

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": claims.Iss,
		"sub": claims.Sub,
		"aud": claims.Aud,
		"exp": claims.Exp,
		"iat": claims.Iat,
		"nbf": claims.Nbf,
		"jti": claims.Jti,
	})

	token, err := tokenObj.SignedString(secret)

	return token, err
}

// ValidateToken expects jwt token, and secret that was used to sign the token
// and returns the claims attached with the token
func ValidateToken(token string, secret []byte) (*jwt.MapClaims, error) {

	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid claims")
}

package dto

type GenerateNonceDTO struct {
	Eth_Addr string `json:"eth_addr" validate:"required,eth_addr"`
}

type GenerateNonceResponseDTO struct {
	Nonce string `json:"nonce"`
}

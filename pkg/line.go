package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

type LineConfiguration struct {
	LineToken string
	LineSecret string
}

type LineSignatureValidator struct {
	secret string
}

func NewLineSignatureValidator(secret string) *LineSignatureValidator {
	return &LineSignatureValidator{
		secret: secret,
	}
}

func (line LineSignatureValidator) Validate(header string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		log.Println(err)
		return false
	}

	hash := hmac.New(sha256.New, []byte(line.secret))
	hash.Write(body)
	if hmac.Equal(hash.Sum(nil), decoded) {
		return true
	}
	return false
}
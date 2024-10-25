package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/spf13/viper"
)

var publicKey *rsa.PublicKey

func InitPublicKey() error {
	publicKeyPEM := viper.GetString("jwt.public-key")
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return fmt.Errorf("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse DER encoded public key: %v", err)
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		publicKey = pub
	default:
		return fmt.Errorf("unknown type of public key")
	}

	return nil
}

func GetPublicKey() *rsa.PublicKey {
	return publicKey
}

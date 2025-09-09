package certmgr

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/pkcs12"
)

// LoadPFX carrega um .pfx/.p12 (bytes) com a senha e retorna um tls.Certificate pronto para usar.
func LoadPFX(pfxData []byte, password string) (tls.Certificate, error) {
	// Extrai todos os blocos PEM
	blocks, err := pkcs12.ToPEM(pfxData, password)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("erro ao converter PFX para PEM: %w", err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	// Monta tls.Certificate direto a partir dos blocos
	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("erro ao montar X509KeyPair: %w", err)
	}

	return cert, nil
}

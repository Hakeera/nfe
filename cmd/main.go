package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/pkcs12"
)

// LoadPFX carrega o arquivo .pfx e retorna um tls.Certificate pronto
func LoadPFX(pfxPath, password string) (tls.Certificate, error) {
	pfxData, err := os.ReadFile(pfxPath)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("erro ao ler PFX: %w", err)
	}

	blocks, err := pkcs12.ToPEM(pfxData, password)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("erro ao converter PFX para PEM: %w", err)
	}

	var certPEM, keyPEM []byte
	for _, b := range blocks {
		if b.Type == "PRIVATE KEY" {
			keyPEM = pem.EncodeToMemory(b)
		} else if b.Type == "CERTIFICATE" {
			certPEM = append(certPEM, pem.EncodeToMemory(b)...)
		}
	}

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("erro ao montar X509KeyPair: %w", err)
	}

	return tlsCert, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	certPath := os.Getenv("CERT_PATH")
	certPassword := os.Getenv("CERT_PASSWORD")
	token := os.Getenv("TOKEN")

	fmt.Println(certPath, certPassword, token)

	// Referência única da nota
	ref := ""

	// Endpoint de homologação
	urlEnvio := "https://homologacao.focusnfe.com.br/v2/nfe?ref=" + ref

	// Carrega o certificado A1
	cert, err := LoadPFX(certPath, certPassword)
	if err != nil {
		panic(err)
	}

	// Cria pool de CAs
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Configuração TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      rootCAs,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Dados da NFe (JSON)
	dadosDaNota := map[string]any{
		"natureza_operacao":           "Remessa",
		"data_emissao":                "2025-09-09T12:00:00",
		"data_entrada_saida":          "2025-09-09T12:00:00",
		"tipo_documento":              "1",
		"finalidade_emissao":          "1",
		"cnpj_emitente":               "55409080000125",
		"nome_emitente":               "ELIANA MARA DE SOUZA ALMEIDA & CIA LTDA",
		"nome_fantasia_emitente":      "Sarracena Uniformes",
		"logradouro_emitente":         "Rua Antônio Blanco",
		"numero_emitente":             "473",
		"bairro_emitente":             "Vila Costa do Sol",
		"municipio_emitente":          "São Carlos",
		"uf_emitente":                 "SP",
		"cep_emitente":                "13566020",
		"inscricao_estadual_emitente": "637056487119",
		"nome_destinatario":           "NF-E EMITIDA EM AMBIENTE DE HOMOLOGACAO - SEM VALOR FISCAL",
		"cpf_destinatario":            "51966818092",
		"telefone_destinatario":       "1196185555",
		"logradouro_destinatario":     "Rua Sao Januario",
		"numero_destinatario":         "99",
		"bairro_destinatario":         "Crespo",
		"municipio_destinatario":      "Manaus",
		"uf_destinatario":             "AM",
		"pais_destinatario":           "Brasil",
		"cep_destinatario":            "69073178",
		"valor_frete":                 "0.0",
		"valor_seguro":                "0",
		"valor_total":                 "47.23",
		"valor_produtos":              "47.23",
		"modalidade_frete":            "0",
		"items": []map[string]any{
			{
				"numero_item":                "1",
				"codigo_produto":             "1232",
				"descricao":                  "Cartões de Visita",
				"cfop":                       "6923",
				"unidade_comercial":          "un",
				"quantidade_comercial":       "100",
				"valor_unitario_comercial":   "0.4723",
				"valor_unitario_tributavel":  "0.4723",
				"unidade_tributavel":         "un",
				"codigo_ncm":                 "49111090",
				"quantidade_tributavel":      "100",
				"valor_bruto":                "47.23",
				"icms_situacao_tributaria":   "400",
				"icms_origem":                "0",
				"pis_situacao_tributaria":    "07",
				"cofins_situacao_tributaria": "07",
			},
		},
	}

	// Converte para JSON
	jsonData, err := json.Marshal(dadosDaNota)
	if err != nil {
		panic(err)
	}

	// Cria a requisição POST
	req, err := http.NewRequest("POST", urlEnvio, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(token, "")

	// Envia a requisição
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Lê resposta
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Código HTTP:", resp.StatusCode)
	fmt.Println("Resposta:", string(body))
}

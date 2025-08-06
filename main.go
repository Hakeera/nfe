package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	clienteID := ObterClienteID()
	fmt.Println("Procurando CNPJ para o Cliente ID:", clienteID)

	// Abre o CSV de Clientes 
	file, err := os.Open("dados/clientes.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Lê o cabeçalho
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Erro ao ler o cabeçalho:", err)
		return
	}

	// Encontra os índices das colunas "id" e "documento"
	idIndex := -1
	cnpjIndex := -1
	nomeIndex := -1
	for i, col := range header {
		switch strings.ToLower(col) {
		case "id":
			idIndex = i
		case "documento":
			cnpjIndex = i
		case "nome":
			nomeIndex = i
		}
	}

	if idIndex == -1 {
		fmt.Println("Coluna 'id' não encontrada")
		return
	}

	if cnpjIndex == -1 {
		fmt.Println("Coluna 'documento' não encontrada")
		return
	}

	if nomeIndex == -1 {
		fmt.Println("Coluna 'documento' não encontrada")
		return
	}
	
	// Procura o clienteID
	for {
		record, err := reader.Read()
		if err != nil {
			break // EOF
		}

		if strings.TrimSpace(record[idIndex]) == clienteID {
			cnpj := strings.TrimSpace(record[cnpjIndex])
			nome := strings.TrimSpace(record[nomeIndex])
			fmt.Println("CNPJ:", cnpj, "Nome:", nome)
			return
		}
	}

	fmt.Println("Cliente ID não encontrado no arquivo.")
}

func ObterClienteID() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o Cliente ID: ")
	clienteID, _ := reader.ReadString('\n')

	return strings.TrimSpace(clienteID)
}

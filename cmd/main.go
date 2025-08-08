package main

import (
	"fmt"
	"log"
	load "nfe/internal/reader"
	"nfe/internal/transformer"
	"nfe/model"
)

func main() {
	// Caminho do CSV
	caminhoclientes := "dados/csv/clientes.csv"
	//caminhopedidos := "dados/csv/pedidos.csv"

	// Carrega os dados do CSV para um slice de Clientes
	clientes, err := load.LoadCSV[model.Cliente](caminhoclientes)
	if err != nil {
		log.Fatalf("Erro ao carregar clientes: %v", err)
	}

	// Carrega os dados do CSV para um slice de Clientes
	//pedidos, err := load.LoadCSV[model.Pedidos](caminhopedidos)
	if err != nil {
		log.Fatalf("Erro ao carregar clientes: %v", err)
	}

	// Imprime os pedidos carregados
	//for i, pedido := range pedidos {
		//fmt.Printf("Pedido %d: %+v\n", i+1, pedido)
		//fmt.Println("Pedido : ", pedido.Situacao////)
	//}
	
	// Imprime os clientes carregados
	for i, cliente := range clientes {
		clienteID := cliente.ID
		if clienteID == "158" {
			cep := cliente.CEP
			 
			endereco, err := transformer.GetEnderecoPorCEP(cep)
			if err != nil {
				fmt.Println("Erro:", err)
				return
			}

			fmt.Printf("Rua: %s, Bairro: %s, Cidade: %s, Estado: %s\n",
			endereco.Logradouro, endereco.Bairro, endereco.Localidade, endereco.UF)
		}
		fmt.Println("ID:", i)
	}
}

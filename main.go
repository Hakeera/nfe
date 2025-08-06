package main

import (
	"fmt"
	"log"
	"nfe/dados"
	"nfe/model"
)

func main() {
	// Caminho do CSV
	caminhoclientes := "dados/csv/clientes.csv"
	caminhopedidos := "dados/csv/pedidos.csv"

	// Carrega os dados do CSV para um slice de Clientes
	clientes, err := dados.LoadCSV[model.Cliente](caminhoclientes)
	if err != nil {
		log.Fatalf("Erro ao carregar clientes: %v", err)
	}

	// Carrega os dados do CSV para um slice de Clientes
	pedidos, err := dados.LoadCSV[model.Pedidos](caminhopedidos)
	if err != nil {
		log.Fatalf("Erro ao carregar clientes: %v", err)
	}
	
	// Imprime os clientes carregados
	for i, cliente := range clientes {
		fmt.Printf("Cliente %d: %+v\n", i+1, cliente)
		fmt.Println("Cliente: ", cliente.Documento)
	}
	// Imprime os pedidos carregados
	for i, pedido := range pedidos {
		fmt.Printf("Pedido %d: %+v\n", i+1, pedido)
		fmt.Println("Pedido : ", pedido.Situacao)
	}
	
}

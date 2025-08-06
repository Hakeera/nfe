// Package model
package model

type NFe struct {

}

type Cliente struct {
	ID        string `csv:"id"`
	Nome      string `csv:"nome"`
	Telefone  string `csv:"telefone"`
	Celular   string `csv:"celular"`
	Documento string `csv:"documento"`
	CNPJ      string `csv:"cnpj"`
	Email     string `csv:"email"`
	Situacao  string `csv:"situacao"`
	IE        string `csv:"ie"`
	CEST      string `csv:"cest"`
	Endereco  string `csv:"endereco"`
	Cidade    string `csv:"cidade"`
	Estado    string `csv:"estado"`
	CEP       string `csv:"cep"`
	Apelido   string `csv:"apelido"`
}

type Pedidos struct {
	PedidoID	string `csv:"pedidoid"`
	ClienteID	string `csv:"clienteid"`
	ValorTotal	string `csv:"valortotal"`
	Situacao	string `csv:"situacao"`
	DataOrcamento	string `csv:"dataorcamento"`
	DataAprovacao	string `csv:"dataaprovacao"`
	DataEntrega	string `csv:"dataentrega"`
	FormaPagamento	string `csv:"formapagamento"`
	NotaFiscal	string `csv:"notafiscal"`
}

type PedidosProdutos struct {

}

type Produtos struct {

}



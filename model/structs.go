// Package model
package model

import "time"

type NFe struct {
	NaturezaOperacao           string   `json:"natureza_operacao"`
	DataEmissao                time.Time `json:"data_emissao"`
	DataEntradaSaida           time.Time `json:"data_entrada_saida"`
	TipoDocumento              int      `json:"tipo_documento"`
	FinalidadeEmissao          int      `json:"finalidade_emissao"`
	CNPJEmitente               string   `json:"cnpj_emitente,omitempty"`
	CPFEmitente                string   `json:"cpf_emitente,omitempty"`
	NomeEmitente               string   `json:"nome_emitente"`
	NomeFantasiaEmitente       string   `json:"nome_fantasia_emitente"`
	LogradouroEmitente         string   `json:"logradouro_emitente"`
	NumeroEmitente             int      `json:"numero_emitente"`
	BairroEmitente             string   `json:"bairro_emitente"`
	MunicipioEmitente          string   `json:"municipio_emitente"`
	UFEmitente                 string   `json:"uf_emitente"`
	CEPEmitente                string   `json:"cep_emitente"`
	InscricaoEstadualEmitente  string   `json:"inscricao_estadual_emitente"`
	NomeDestinatario           string   `json:"nome_destinatario"`
	CPFDestinatario            string   `json:"cpf_destinatario,omitempty"`
	InscricaoEstadualDestinatario *string `json:"inscricao_estadual_destinatario,omitempty"`
	TelefoneDestinatario       string   `json:"telefone_destinatario"`
	LogradouroDestinatario     string   `json:"logradouro_destinatario"`
	NumeroDestinatario         int      `json:"numero_destinatario"`
	BairroDestinatario         string   `json:"bairro_destinatario"`
	MunicipioDestinatario      string   `json:"municipio_destinatario"`
	UFDestinatario             string   `json:"uf_destinatario"`
	PaisDestinatario           string   `json:"pais_destinatario"`
	CEPDestinatario            string   `json:"cep_destinatario"`
	ValorFrete                 float64  `json:"valor_frete"`
	ValorSeguro                float64  `json:"valor_seguro"`
	ValorTotal                 float64  `json:"valor_total"`
	ValorProdutos              float64  `json:"valor_produtos"`
	ModalidadeFrete            int      `json:"modalidade_frete"`
	Items                      []NFeItem `json:"items"`
}

type NFeItem struct {
    NumeroItem                 int     `json:"numero_item"`
    CodigoProduto              int     `json:"codigo_produto"`
    Descricao                  string  `json:"descricao"`
    CFOP                       int     `json:"cfop"`
    UnidadeComercial           string  `json:"unidade_comercial"`
    QuantidadeComercial        float64 `json:"quantidade_comercial"`
    ValorUnitarioComercial     float64 `json:"valor_unitario_comercial"`
    ValorUnitarioTributavel    float64 `json:"valor_unitario_tributavel"`
    UnidadeTributavel          string  `json:"unidade_tributavel"`
    CodigoNCM                  int     `json:"codigo_ncm"`
    QuantidadeTributavel       float64 `json:"quantidade_tributavel"`
    ValorBruto                 float64 `json:"valor_bruto"`
    ICMSSituacaoTributaria     int     `json:"icms_situacao_tributaria"`
    ICMSOrigem                 int     `json:"icms_origem"`
    PISSituacaoTributaria      string  `json:"pis_situacao_tributaria"`
    CofinsSituacaoTributaria   string  `json:"cofins_situacao_tributaria"`
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
	PedidoID	string `csv:"pedidoid"`
	ClienteID	string `csv:"clienteid"`
	ProdutoID	string `csv:"produtoid"`
	Nome		string `csv:"nome"`
	QtdProduto	int `csv:"qtdproduto"`
	PrecoProduto	int `csv:"precoproduto"`
	Total		int `csv:"total"`
}

type Produtos struct {
	ProdutoID	string `csv:"produtoid"`
	Modelo		string `csv:"modelo"`
	Tecido		string `csv:"tecido"`
	Cor		string `csv:"cor"`
	Cliente		string `csv:"cliente"`
	Tamanho		string `csv:"tamanho"`
	Linha		string `csv:"linha"`
	Situacao	string `csv:"situacao"`
	Tabela1		int    `csv:"tabela1"`
	Tabela2		int    `csv:"tabela2"`
	Descricao	string `csv:"descricao"`
	Nome		string `csv:"nome"`
	FichaTecnica	string `csv:"fichatecnica"`
}

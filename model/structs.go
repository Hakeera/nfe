// Package model
package model

type cliente struct {
	id	  string	 `csv:"id"` 
	nome	  string	 `csv:"nome"`
	telefone  string	 `csv:"telefone"`
	celular	  string	 `csv:"celular"`
	documento string	 `csv:"documento"`
	cnpj	  string	 `csv:"cnpj"`
	email     string	 `csv:"email"`
	situacao  string	 `csv:"situacao"`
	ie	  string	 `csv:"ie"`
	cest	  string	 `csv:"cest"`
	endereco  string	 `csv:"endereco"`
	cidade	  string	 `csv:"cidade"`
	estado	  string	 `csv:"estado"`
	cep	  string	 `csv:"cep"`
	apelido	  string	 `csv:"apelido"`
}

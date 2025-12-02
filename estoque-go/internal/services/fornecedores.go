package services

import "fmt"

type FornecedorService interface {
	GetInfo() string
}

type Fornecedor struct {
	CNPJ    string
	Contato string
	Cidade  string
}

func (f Fornecedor) GetInfo() string {
	return fmt.Sprintf("CNPJ: %s | Contato: %s | Cidade: %s", f.CNPJ, f.Contato, f.Cidade)
}

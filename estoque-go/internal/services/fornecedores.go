package services

import "fmt"

type InfoService interface {
	GetInfo() string
}

type Disponibilidade interface {
	VerificarDisponibilidade(quantidadeSolicitada int, quantidadeDisponivel int) bool
}

type FornecedorService interface {
	InfoService
	Disponibilidade
}

type Fornecedor struct {
	CNPJ    string
	Contato string
	Cidade  string
}

func (f Fornecedor) GetInfo() string {
	return fmt.Sprintf("CNPJ: %s | Contato: %s | Cidade: %s", f.CNPJ, f.Contato, f.Cidade)
}

func (f Fornecedor) VerificarDisponibilidade(quantidadeSolicitada int, quantidadeDisponivel int) bool {
	return quantidadeSolicitada <= quantidadeDisponivel
}

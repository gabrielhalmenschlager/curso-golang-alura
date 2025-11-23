package models

import "time"

type Categoria string

const (
	BRONZE Categoria = "BRONZE"
	PRATA  Categoria = "PRATA"
	OURO   Categoria = "OURO"
)

type Jogo struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Nome         string    `json:"nome"`
	Descricao    string    `json:"descricao"`
	Categoria    Categoria `json:"categoria"`
	Disponivel   bool      `json:"disponivel" gorm:"default:true"`
	DataCadastro time.Time `json:"data_cadastro" gorm:"autoCreateTime"`

	ImagemURL string `json:"imagem_url"`
}

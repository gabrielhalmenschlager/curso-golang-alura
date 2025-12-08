package models

type Categoria struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Nome      string `json:"nome"`
	Codigo    string `gorm:"unique" json:"codigo"`
	Descricao string `json:"descricao"`
}

func (Categoria) TableName() string {
	return "categorias"
}

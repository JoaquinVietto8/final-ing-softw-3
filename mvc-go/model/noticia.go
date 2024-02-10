package model

type Noticia struct {
	Id        int    `gorm:"primaryKey"`
	Contenido string `gorm:"type:varchar(350);not null"`
	Fecha     string `gorm:"type:date;not null"`
	Imagen    string `gorm:"type:varchar(400);not null"`
}

type Noticias []Noticia

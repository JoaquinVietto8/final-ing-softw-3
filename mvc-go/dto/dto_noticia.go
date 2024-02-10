package dto

type NoticiaDto struct {
	Id_noticia int    `json:"id_noticia"`
	Contenido  string `json:"contenido"`
	Fecha      string `json:"fecha"`
	Imagen     string `json:"imagen"`
}

type NoticiasDto []NoticiaDto

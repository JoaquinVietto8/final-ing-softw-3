package services

import (
	noticiaCliente "mvc-go/clients/noticia"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
)

type noticiaService struct{}

type noticiaServiceInterface interface {
	GetNoticias() (dto.NoticiasDto, e.ApiError)
	InsertNoticia(noticiaDto dto.NoticiaDto) (dto.NoticiaDto, e.ApiError)
}

var (
	NoticiaService noticiaServiceInterface
)

func init() {
	NoticiaService = &noticiaService{}
}

func (s *noticiaService) GetNoticias() (dto.NoticiasDto, e.ApiError) {

	var noticias model.Noticias = noticiaCliente.GetNoticias()
	var noticiasDto dto.NoticiasDto

	for _, noticia := range noticias {
		var noticiaDto dto.NoticiaDto
		noticiaDto.Id_noticia = noticia.Id
		noticiaDto.Contenido = noticia.Contenido
		noticiaDto.Imagen = noticia.Imagen
		noticiaDto.Fecha = noticia.Fecha

		noticiasDto = append(noticiasDto, noticiaDto)
	}

	return noticiasDto, nil
}

func (s *noticiaService) InsertNoticia(noticiaDto dto.NoticiaDto) (dto.NoticiaDto, e.ApiError) {

	var noticia model.Noticia

	noticia.Fecha = time.Now().Format("2006.01.02")
	noticia.Contenido = noticiaDto.Contenido
	noticia.Imagen = noticiaDto.Imagen

	noticiaCliente.InsertNoticia(noticia)

	return noticiaDto, nil
}

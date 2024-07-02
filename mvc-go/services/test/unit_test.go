package test

import (
	e "final-ing-softw-3/mvc-go/utils/errors"
	"fmt"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Se crea un cliente mock para pruebas
type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

type noticiaService struct {
	noticiaCliente *MockClient
}

func (s *noticiaService) GetNoticias() ([]model.Noticia, e.ApiError) {
	// Noticias de ejemplo
	return []model.Noticia{
		{Id: 1, Contenido: "Contenido de la noticia 1", Imagen: "imagen1.jpg", Fecha: "2024-02-11"},
		{Id: 2, Contenido: "Contenido de la noticia 2", Imagen: "imagen2.jpg", Fecha: "2024-02-12"},
	}, nil
}

func TestGetNoticias(t *testing.T) {

	// Crea una instancia del servicio con el cliente mock
	service := &noticiaService{noticiaCliente: NewMockClient()}

	// Llama a la función que se está probando
	noticias, err := service.GetNoticias()

	// Verifica que no haya error
	assert.NoError(t, err, "Ocurrió un error")

	// Verifica que se devuelvan las noticias esperadas
	var noticiasDto dto.NoticiasDto
	for _, noticia := range noticias {
		var noticiaDto dto.NoticiaDto

		noticiaDto.Id_noticia = noticia.Id
		noticiaDto.Contenido = noticia.Contenido
		noticiaDto.Imagen = noticia.Imagen
		noticiaDto.Fecha = noticia.Fecha

		noticiasDto = append(noticiasDto, noticiaDto)
	}

	expectedNoticiasDto := dto.NoticiasDto{
		{Id_noticia: 1, Contenido: "Contenido de la noticia 1", Imagen: "imagen1.jpg", Fecha: "2024-02-11"},
		{Id_noticia: 2, Contenido: "Contenido de la noticia 2", Imagen: "imagen2.jpg", Fecha: "2024-02-12"},
	}
	assert.Equal(t, expectedNoticiasDto, noticiasDto, "Las noticias devueltas no coinciden")
}

// Se simula un error
/*	expectedNoticiasDto := dto.NoticiasDto{
		{Id_noticia: 1, Contenido: "Contenido de la noticia 1", Imagen: "imagen1.jpg", Fecha: "2024-02-11"},
		{Id_noticia: 5, Contenido: "Contenido de la noticia 5", Imagen: "imagen5.jpg", Fecha: "2024-02-25"},
	}
	assert.Equal(t, expectedNoticiasDto, noticiasDto, "Las noticias devueltas no coinciden") */

func (s *noticiaService) InsertNoticia(noticia model.Noticia) e.ApiError {
	fmt.Printf("Noticia insertada: %+v\n", noticia)
	return nil
}

func TestInsertNoticia(t *testing.T) {
	// Crea una instancia del servicio con el cliente mock
	service := &noticiaService{noticiaCliente: NewMockClient()}

	// Se define una noticia de prueba
	noticiaDto := dto.NoticiaDto{
		Contenido: "Contenido de la noticia",
		Imagen:    "imagen.jpg",
	}

	// Se convierten dto.NoticiaDto a model.Noticia
	noticia := model.Noticia{
		Contenido: noticiaDto.Contenido,
		Imagen:    noticiaDto.Imagen,
	}

	// Llama a la función que se está probando
	err := service.InsertNoticia(noticia)

	// Verifica que no haya error
	assert.NoError(t, err, "Ocurrio un error")
}

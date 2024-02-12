package services

import (
	"fmt"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Se crea un cliente mock para pruebas
type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (m *MockClient) GetNoticias() (model.Noticia, e.ApiError) {
	return model.Noticia{
		{Id: 1, Contenido: "Contenido de la noticia 1", Imagen: "imagen1.jpg", Fecha: "2024-02-11"},
		{Id: 2, Contenido: "Contenido de la noticia 2", Imagen: "imagen2.jpg", Fecha: "2024-02-12"},
	}, nil
}

func TestGetNoticias(t *testing.T) {

	// Crea una instancia del servicio con el cliente mock
	service := &noticiaService{noticiaCliente: NewMockClient()}

	// Llama a la función que se está probando
	noticiasDto, err := service.GetNoticias()

	// Verifica que no haya error
	assert.NoError(t, err, "Ocurrio un error")

	// Verifica que se devuelvan las noticias esperadas
	expectedNoticiasDto := dto.NoticiasDto{
		{Id_noticia: 1, Contenido: "Contenido de la noticia 1", Imagen: "imagen1.jpg", Fecha: "2024-02-11"},
		{Id_noticia: 2, Contenido: "Contenido de la noticia 2", Imagen: "imagen2.jpg", Fecha: "2024-02-12"},
	}
	assert.Equal(t, expectedNoticiasDto, noticiasDto, "Las noticias devueltas no coinciden con las esperadas")
}

func (m *MockClient) InsertNoticia(noticia model.Noticia) e.ApiError {
	fmt.Printf("Noticia insertada en el cliente mock: %+v\n", noticia)
	return nil
}

func TestInsertNoticia(t *testing.T) {
	// Crea una instancia del servicio con el cliente mock
	service := &noticiaService{noticiaCliente: NewMockClient()}

	// Define una noticia de prueba
	noticiaDto := dto.NoticiaDto{
		Contenido: "Contenido de la noticia",
		Imagen:    "imagen.jpg",
	}

	// Llama a la función que se está probando
	_, err := service.InsertNoticia(noticiaDto)

	// Verifica que no haya error
	assert.NoError(t, err, "Ocurrio un error")
}

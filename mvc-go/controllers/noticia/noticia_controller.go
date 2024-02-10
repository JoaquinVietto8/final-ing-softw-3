package noticiaController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetNoticias(c *gin.Context) {
	var noticiasDto dto.NoticiasDto
	noticiasDto, err := service.NoticiaService.GetNoticias()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, noticiasDto)
}

func InsertNoticia(c *gin.Context) {

	var noticiaDto dto.NoticiaDto
	err := c.BindJSON(&noticiaDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	noticiaDto, er := service.NoticiaService.InsertNoticia(noticiaDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, noticiaDto)
}

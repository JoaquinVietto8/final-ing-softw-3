package app

import (
	noticiaController "mvc-go/controllers/noticia"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Noticia Mapping
	router.GET("/inicio", noticiaController.GetNoticias)
	router.POST("/nueva-noticia", noticiaController.InsertNoticia)

	log.Info("Finishing mappings configurations")
}

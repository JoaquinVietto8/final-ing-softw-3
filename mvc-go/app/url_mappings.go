package app

import (
	"final-ing-softw-3/mvc-go/clients/noticia"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func MapUrls(router *mux.Router) {
	// Noticia Mapping
	router.HandleFunc("/inicio", noticia.GetNoticias).Methods("GET")
	router.HandleFunc("/nueva-noticia", noticia.InsertNoticia).Methods("POST")

	logrus.Info("Finishing mappings configurations")
}

/*
func mapUrls() {
	// Noticia Mapping
	router.GET("/inicio", noticiaController.GetNoticias)
	router.POST("/nueva-noticia", noticiaController.InsertNoticia)

	log.Info("Finishing mappings configurations")
}
*/

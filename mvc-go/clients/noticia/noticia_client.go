package noticia

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetNoticias() model.Noticias {
	var noticias model.Noticias
	Db.Find(&noticias)
	log.Debug("Noticias: ", noticias)

	return noticias
}

func InsertNoticia(noticia model.Noticia) model.Noticia {

	result := Db.Create(&noticia)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Noticia Created: ", noticia.Id)
	return noticia
}

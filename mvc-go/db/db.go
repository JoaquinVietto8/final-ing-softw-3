package db

import (
	noticiaClient "mvc-go/clients/noticia"
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "diario"
	DBUser := "root"
	DBPass := "root"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "database"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	noticiaClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	//agregar el resto de los objetos del model
	db.AutoMigrate(&model.Noticia{})

	log.Info("Finishing Migration Database Tables")
}

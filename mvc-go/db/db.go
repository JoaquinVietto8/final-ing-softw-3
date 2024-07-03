package db

import (
	"fmt"
	noticiaClient "mvc-go/clients/noticia"
	"mvc-go/model"
	"os"

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
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	// ------------------------

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True", DBUser, DBPass, DBHost, DBName)
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Error("Connection Failed to Open: ", err)
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

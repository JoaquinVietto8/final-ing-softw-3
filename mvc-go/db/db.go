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
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")
	//dbUser := os.Getenv("DB_USER")
	//dbPass := os.Getenv("DB_PASS")
	// ------------------------

	db, err = gorm.Open("mysql", "umqefiasb4eerhdc:agyJvR3b86TriD2SciLD@tcp(bavlh4vaq9znl1j1mh7v-mysql.services.clever-cloud.com:3306)/bavlh4vaq9znl1j1mh7v?charset=utf8&parseTime=True")

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

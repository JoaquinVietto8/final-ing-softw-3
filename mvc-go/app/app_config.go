package app

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.Info("Starting logger system")

	/*
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	*/
}

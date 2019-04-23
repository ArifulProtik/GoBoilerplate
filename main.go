package main

import (
	"fmt"
	"net/http"
	"os"

	logger "github.com/anikhasibul/log"
	"github.com/gorilla/handlers"
	"github.com/ArifulProtik/GoBoilerplate/api/handler"
)

var log = logger.New(os.Stdout)

func main() {
	var addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	if os.Getenv("PORT") == "" {
		addr = "127.0.0.1:8080"
	}
	log.Info("Server starting on", addr)
	err := http.ListenAndServe(addr, handlers.CompressHandler(router.Router))
	if err != nil {
		log.Fatal(err)
	}
}

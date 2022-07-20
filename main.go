package main

import (
	"flightApi/src/config"
	"flightApi/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfigApp()
	fmt.Println("Running")
	fmt.Println("Try to read a DB Uri: ", config.StringConnDb)
	r := router.CreateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

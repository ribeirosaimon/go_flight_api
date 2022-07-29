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
	fmt.Println(fmt.Sprintf("Running in port: %d", config.Port))
	r := router.CreateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

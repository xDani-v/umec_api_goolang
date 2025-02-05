package main

import (
	"log"
	"net/http"

	"github.com/xDani-v/umec_api_goolang/data"
	"github.com/xDani-v/umec_api_goolang/routes"
)

func main() {
	data.Connect()
	rutas := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", rutas))
	//imprimri mensaje consola de funcionamiento
	log.Println("Server running on port: http://localhost:8080/api/")
}


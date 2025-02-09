package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/xDani-v/umec_api_goolang/data"
	"github.com/xDani-v/umec_api_goolang/models"
	"github.com/xDani-v/umec_api_goolang/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	data.Connect()
	data.DB.AutoMigrate(&models.Rol{}, &models.Usuario{}, &models.Parametros{}, &models.ParametrosValor{})
	rutas := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", rutas))
	log.Println("Server running on port: http://localhost:8080/api/")
}

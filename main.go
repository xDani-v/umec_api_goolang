package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"  // Add this import
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
	data.DB.AutoMigrate(&models.Rol{}, &models.Usuario{}, &models.Parametros{}, &models.ParametrosValor{}, &models.Especialidades{},&models.Especialidades{}, &models.Cita{}, &models.Menu{}, &models.Funcionalidad{}, &models.RolesFuncionalidad{})
	rutas := routes.InitRouter()

	// CORS configuration
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// Apply CORS middleware
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(rutas)
	
	log.Println("Server running on port: http://localhost:8080/api/")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
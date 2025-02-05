package routes

import (
	"github.com/gorilla/mux"
	"github.com/xDani-v/umec_api_goolang/controllers"
)

func InitRouter() *mux.Router{
	 rutas := mux.NewRouter()
	 api := rutas.PathPrefix("/api").Subrouter()

	 api.HandleFunc("/test", controllers.GetIniciar).Methods("GET")

	 return rutas
 }
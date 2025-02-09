package routes

import (
	"github.com/gorilla/mux"
	"github.com/xDani-v/umec_api_goolang/controllers"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()

	api.HandleFunc("/test", controllers.GetIniciar).Methods("GET")

	apiRoles := api.PathPrefix("/roles").Subrouter()
	apiRoles.HandleFunc("", controllers.GetRoles).Methods("GET")
	apiRoles.HandleFunc("", controllers.CreateRol).Methods("POST")
	apiRoles.HandleFunc("", controllers.UpdateRol).Methods("PUT")
	apiRoles.HandleFunc("/{id}", controllers.DeleteRol).Methods("DELETE")
	apiRoles.HandleFunc("/{id}", controllers.GetRol).Methods("GET")

	apiUsuarios := api.PathPrefix("/usuario").Subrouter()
	apiUsuarios.HandleFunc("", controllers.GetUsuarios).Methods("GET")
	apiUsuarios.HandleFunc("", controllers.Register).Methods("POST")
	apiUsuarios.HandleFunc("/login", controllers.Login).Methods("POST")
	apiUsuarios.HandleFunc("/{id}", controllers.GetUsuario).Methods("GET")
	apiUsuarios.HandleFunc("/{id}", controllers.UpdateUsuario).Methods("PUT")
	apiUsuarios.HandleFunc("/{id}", controllers.DeleteUsuario).Methods("DELETE")

	return rutas
}

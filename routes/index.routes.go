package routes

import (
	"github.com/gorilla/mux"
	"github.com/xDani-v/umec_api_goolang/controllers"
	"github.com/xDani-v/umec_api_goolang/middleware"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()

	api.HandleFunc("/test", controllers.GetIniciar).Methods("GET")

	apiRoles := api.PathPrefix("/roles").Subrouter()
	apiRoles.Use(middleware.AuthMiddleware)
	apiRoles.HandleFunc("", controllers.GetRoles).Methods("GET")
	apiRoles.HandleFunc("", controllers.CreateRol).Methods("POST")
	apiRoles.HandleFunc("", controllers.UpdateRol).Methods("PUT")
	apiRoles.HandleFunc("/{id}", controllers.DeleteRol).Methods("DELETE")
	apiRoles.HandleFunc("/{id}", controllers.GetRol).Methods("GET")

	// Rutas públicas para usuarios
	apiUsuarios := api.PathPrefix("/usuario").Subrouter()
	apiUsuarios.HandleFunc("/login", controllers.Login).Methods("POST")
	apiUsuarios.HandleFunc("", controllers.Register).Methods("POST")

	// Rutas protegidas para usuarios
	apiUsuariosProtected := api.PathPrefix("/usuario").Subrouter()
	apiUsuariosProtected.Use(middleware.AuthMiddleware) // Aplicar middleware de autenticación
	apiUsuariosProtected.HandleFunc("", controllers.GetUsuarios).Methods("GET")
	apiUsuariosProtected.HandleFunc("/{id}", controllers.GetUsuario).Methods("GET")
	apiUsuariosProtected.HandleFunc("/{id}", controllers.UpdateUsuario).Methods("PUT")
	apiUsuariosProtected.HandleFunc("/{id}", controllers.DeleteUsuario).Methods("DELETE")

	//Parametros
	apiParametros := api.PathPrefix("/parametros").Subrouter()
	apiParametros.Use(middleware.AuthMiddleware)
	apiParametros.HandleFunc("", controllers.GetParametros).Methods("GET")
	apiParametros.HandleFunc("", controllers.CreateParametro).Methods("POST")
	apiParametros.HandleFunc("", controllers.UpdateParametro).Methods("PUT")
	apiParametros.HandleFunc("/{id}", controllers.DeleteParametro).Methods("DELETE")

	//parametros valores
	apiParametrosValores := api.PathPrefix("/parametrovalor").Subrouter()
	apiParametrosValores.Use(middleware.AuthMiddleware)
	apiParametrosValores.HandleFunc("", controllers.GetParametrosValor).Methods("GET")
	apiParametrosValores.HandleFunc("/{id}", controllers.GetParametroValor).Methods("GET")
	apiParametrosValores.HandleFunc("", controllers.CreateParametroValor).Methods("POST")
	apiParametrosValores.HandleFunc("", controllers.UpdateParametroValor).Methods("PUT")
	apiParametrosValores.HandleFunc("/{id}", controllers.DeleteParametroValor).Methods("DELETE")

	apiEmail := api.PathPrefix("/email").Subrouter()
	apiEmail.Use(middleware.AuthMiddleware)
	apiEmail.HandleFunc("/otp", controllers.EnviarCodigoVerificacion).Methods("POST")

	return rutas
}

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/xDani-v/umec_api_goolang/data"
	"github.com/xDani-v/umec_api_goolang/models"
	"github.com/xDani-v/umec_api_goolang/utils"
)

func GetParametros(w http.ResponseWriter, r *http.Request) {
	var objetos []models.Parametros
	data.DB.Find(&objetos)

	respuesta := utils.ResponseMsg{
		Msg:  "Lista de Parametros",
		Data: objetos,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func GetParametro(w http.ResponseWriter, r *http.Request) {
	objeto := models.Parametros{}
	data.DB.First(&objeto, r.URL.Query().Get("id"))

	respuesta := utils.ResponseMsg{
		Msg:    "Parametro",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func CreateParametro(w http.ResponseWriter, r *http.Request) {
	objeto := models.Parametros{}
	json.NewDecoder(r.Body).Decode(&objeto)
	data.DB.Create(&objeto)

	respuesta := utils.ResponseMsg{
		Msg:    "Parametro creado",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func UpdateParametro(w http.ResponseWriter, r *http.Request) {
	objeto := models.Parametros{}
	json.NewDecoder(r.Body).Decode(&objeto)
	data.DB.Save(&objeto)

	respuesta := utils.ResponseMsg{
		Msg:    "Parametro actualizado",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func DeleteParametro(w http.ResponseWriter, r *http.Request) {
	objeto := models.Parametros{}
	data.DB.First(&objeto, r.URL.Query().Get("id"))
	data.DB.Delete(&objeto)

	respuesta := utils.ResponseMsg{
		Msg:    "Parametro eliminado",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

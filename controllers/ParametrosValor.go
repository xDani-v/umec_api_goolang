package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/xDani-v/umec_api_goolang/data"
	"github.com/xDani-v/umec_api_goolang/models"
	"github.com/xDani-v/umec_api_goolang/utils"
)

func GetParametrosValor(w http.ResponseWriter, r *http.Request) {
	var objetos []models.ParametrosValor
	data.DB.Find(&objetos)

	respuesta := utils.ResponseMsg{
		Msg:  "Lista de Parametros Valor",
		Data: objetos,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func GetParametroValor(w http.ResponseWriter, r *http.Request) {
	objeto := models.ParametrosValor{}
	data.DB.First(&objeto, r.URL.Query().Get("id"))

	respuesta := utils.ResponseMsg{
		Msg:    "Parametro Valor",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func CreateParametroValor(w http.ResponseWriter, r *http.Request) {
	objeto := models.ParametrosValor{}
	json.NewDecoder(r.Body).Decode(&objeto)
	data.DB.Create(&objeto)

	respuesta := utils.ResponseMsg{
		Msg:    "Parametro Valor creado",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func UpdateParametroValor(w http.ResponseWriter, r *http.Request) {
	objeto := models.ParametrosValor{}
	json.NewDecoder(r.Body).Decode(&objeto)
	data.DB.Save(&objeto)

	respuesta := utils.ResponseMsg{
		Msg:    "Parametro Valor actualizado",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func DeleteParametroValor(w http.ResponseWriter, r *http.Request) {
	objeto := models.ParametrosValor{}
	data.DB.First(&objeto, r.URL.Query().Get("id"))
	data.DB.Delete(&objeto)

	respuesta := utils.ResponseMsg{
		Msg:    "Parametro Valor eliminado",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

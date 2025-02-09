package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/xDani-v/umec_api_goolang/data"
	"github.com/xDani-v/umec_api_goolang/models"
	"github.com/xDani-v/umec_api_goolang/utils"
)

const contentType = "Content-Type"
const applicationJSON = "application/json"

func GetRoles(w http.ResponseWriter, r *http.Request) {
	var roles []models.Rol
	data.DB.Find(&roles)

	respuesta := utils.ResponseMsg{
		Msg:  "Roles",
		Data: roles,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func GetRol(w http.ResponseWriter, r *http.Request) {
	rol := models.Rol{}
	data.DB.First(&rol, r.URL.Query().Get("id"))

	respuesta := utils.ResponseMsg{
		Msg:    "Rol encontrado",
		Data:   rol,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func CreateRol(w http.ResponseWriter, r *http.Request) {
	rol := models.Rol{}
	json.NewDecoder(r.Body).Decode(&rol)
	data.DB.Create(&rol)

	respuesta := utils.ResponseMsg{
		Msg:    "Rol creado",
		Data:   rol,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func UpdateRol(w http.ResponseWriter, r *http.Request) {
	rol := models.Rol{}
	json.NewDecoder(r.Body).Decode(&rol)
	data.DB.Save(&rol)

	respuesta := utils.ResponseMsg{
		Msg:    "Rol actualizado",
		Data:   rol,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func DeleteRol(w http.ResponseWriter, r *http.Request) {
	rol := models.Rol{}
	data.DB.First(&rol, r.URL.Query().Get("id"))
	data.DB.Delete(&rol)

	respuesta := utils.ResponseMsg{
		Msg:    "Rol eliminado",
		Data:   rol,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

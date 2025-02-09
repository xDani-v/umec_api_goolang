package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xDani-v/umec_api_goolang/data"
	"github.com/xDani-v/umec_api_goolang/models"
	"github.com/xDani-v/umec_api_goolang/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	var objetos []models.Usuario
	data.DB.Find(&objetos)

	respuesta := utils.ResponseMsg{
		Msg:  "Lista de Usuarios",
		Data: objetos,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	objeto := models.Usuario{}
	data.DB.First(&objeto, r.URL.Query().Get("id"))

	respuesta := utils.ResponseMsg{
		Msg:    "Usuario",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Cifrar contraseña antes de guardarla
	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al cifrar la contraseña", http.StatusInternalServerError)
		return
	}
	usuario.Password = string(hash)

	if err := data.DB.Create(&usuario).Error; err != nil {
		http.Error(w, "Error al registrar usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado con éxito"})
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	objeto := models.Usuario{}
	json.NewDecoder(r.Body).Decode(&objeto)
	data.DB.Save(&objeto)

	respuesta := utils.ResponseMsg{
		Msg:    "Usuario actualizado",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	objeto := models.Usuario{}
	data.DB.First(&objeto, r.URL.Query().Get("id"))
	data.DB.Delete(&objeto)

	respuesta := utils.ResponseMsg{
		Msg:    "Usuario eliminado",
		Data:   objeto,
		Status: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credenciales struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credenciales); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	if err := data.DB.Where("correo = ?", credenciales.Correo).First(&usuario).Error; err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
		return
	}

	// Verificar la contraseña
	if err := models.VerificarPassword(usuario.Password, credenciales.Password); err != nil {
		http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
		return
	}

	// Generar el token con duración de 2 horas
	token, err := utils.GenerarToken(usuario.ID, usuario.Correo, 2*time.Hour)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	// Devolver el token en la respuesta
	respuesta := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

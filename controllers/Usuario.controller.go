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
        respuesta := utils.ResponseMsg{
            Msg:    "Datos inválidos",
            Status: http.StatusBadRequest,
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(respuesta)
        return
    }

    // Create user in database - BeforeSave hook will handle password hashing
    if err := data.DB.Create(&usuario).Error; err != nil {
        respuesta := utils.ResponseMsg{
            Msg:    "Error al registrar usuario",
            Status: http.StatusInternalServerError,
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(respuesta)
        return
    }

    respuesta := utils.ResponseMsg{
        Msg:    "Usuario registrado exitosamente",
        Data:   usuario.ID,
        Status: http.StatusCreated,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respuesta)
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
        utils.SendError(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    var usuario models.Usuario
    if err := data.DB.Where("correo = ? AND estado = ?", credenciales.Correo, true).First(&usuario).Error; err != nil {
        utils.SendError(w, "Usuario no encontrado o inactivo", http.StatusUnauthorized)
        return
    }

    // Verify password using bcrypt's CompareHashAndPassword directly
    if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(credenciales.Password)); err != nil {
        utils.SendError(w, "Credenciales incorrectas", http.StatusUnauthorized)
        return
    }

    // Generate token with user data
    token, err := utils.GenerarToken(usuario.ID, usuario.Correo, 2*time.Hour)
    if err != nil {
        utils.SendError(w, "Error al generar el token", http.StatusInternalServerError)
        return
    }

    // Return success response with token and user data
    respuesta := utils.ResponseMsg{
        Msg: "Login exitoso",
        Data: map[string]interface{}{
            "token": token,
            "usuario": map[string]interface{}{
                "id":      usuario.ID,
                "correo":  usuario.Correo,
                "nombres": usuario.Nombres,
                "id_rol":  usuario.Id_rol,
            },
        },
        Status: http.StatusOK,
    }

    w.Header().Set(contentType, applicationJSON)
    json.NewEncoder(w).Encode(respuesta)
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xDani-v/umec_api_goolang/utils"
)

type VerificacionRequest struct {
	Email string `json:"correo"`
}

func EnviarCodigoVerificacion(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud
	var req VerificacionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Generar el código de verificación
	codigo := utils.GenerarCodigoVerificacion()

	// Leer la plantilla HTML
	plantilla, err := utils.LeerPlantillaHTML("templates/verificacion.html")
	if err != nil {
		http.Error(w, "Error al leer la plantilla HTML", http.StatusInternalServerError)
		return
	}

	// Reemplazar el marcador de posición con el código de verificación
	mensaje := fmt.Sprintf(plantilla, codigo)

	// Enviar el correo electrónico
	destinatario := req.Email
	asunto := "Código de Verificación"
	if err := utils.EnviarCorreo(destinatario, asunto, mensaje); err != nil {
		http.Error(w, "Error al enviar el correo", http.StatusInternalServerError)
		return
	}

	response := utils.ResponseMsg{
		Msg:    "Correo enviado",
		Data:   codigo,
		Status: http.StatusOK,
	}

	// Devolver la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

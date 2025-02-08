package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/xDani-v/umec_api_goolang/utils"
)

func GetIniciar(w http.ResponseWriter, r *http.Request) {
	res := utils.ResponseMsg{
		Msg:    "API is up and running",
		Data:   "Welcome to UMEC API",
		Status: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

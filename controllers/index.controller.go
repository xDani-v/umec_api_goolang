package controllers

import (
	"encoding/json"
	"net/http"
)

type ResponseMsg struct {
	Msg string `json:"msg"`
	Status int `json:"code"`
}

func GetIniciar(w http.ResponseWriter, r *http.Request) {
	 res := ResponseMsg{
		Msg: "API is up and running",
		Status: 200,
	 }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}


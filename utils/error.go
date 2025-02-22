package utils

import (
    "encoding/json"
    "net/http"
)

func SendError(w http.ResponseWriter, message string, status int) {
    response := ResponseMsg{
        Msg:    message,
        Status: status,
        Data:   nil,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)
}
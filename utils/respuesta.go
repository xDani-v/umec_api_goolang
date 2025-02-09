package utils

import (
	"io/ioutil"
	"math/rand"
	"time"
)

type ResponseMsg struct {
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Status int         `json:"code"`
}

func GenerarCodigoVerificacion() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000)
}

func LeerPlantillaHTML(ruta string) (string, error) {
	contenido, err := ioutil.ReadFile(ruta)
	if err != nil {
		return "", err
	}
	return string(contenido), nil
}

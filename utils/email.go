package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// Configuración del servidor SMTP
var (
	SMTP_HOST     string
	SMTP_PORT     int
	SMTP_USER     string
	SMTP_PASSWORD string
)

// EnviarCorreo envía un email con el asunto y mensaje al destinatario
func EnviarCorreo(destinatario, asunto, mensaje string) error {
	if err := godotenv.Load(); err != nil {
		log.Println("Error al cargar el archivo .env:", err)
	}

	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT = 587 // Puedes cambiar esto si necesitas leerlo desde el archivo .env
	SMTP_USER = os.Getenv("SMTP_USER")
	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", SMTP_USER)
	m.SetHeader("To", destinatario)
	m.SetHeader("Subject", asunto)
	m.SetBody("text/html", mensaje)

	d := gomail.NewDialer(SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASSWORD)
	d.TLSConfig = nil

	if err := d.DialAndSend(m); err != nil {
		log.Println("Error al enviar el correo:", err)
		return err
	}
	return nil
}

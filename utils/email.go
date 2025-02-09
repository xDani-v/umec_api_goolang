package utils

import (
	"log"

	"gopkg.in/gomail.v2"
)

// Configuración del servidor SMTP
const (
	SMTP_HOST     = "smtp.gmail.com"      // Servidor SMTP (Ejemplo: Gmail)
	SMTP_PORT     = 587                   // Puerto SMTP (465 para SSL, 587 para TLS)
	SMTP_USER     = "tu_correo@gmail.com" // Tu correo
	SMTP_PASSWORD = "tu_contraseña"       // Contraseña o App Password
)

// EnviarCorreo envía un email con el asunto y mensaje al destinatario
func EnviarCorreo(destinatario, asunto, mensaje string) error {
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
	log.Println("Correo enviado correctamente a:", destinatario)
	return nil
}

package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model

	ID                 uint64 `json:"ID" gorm:"primary_key;autoIncrement"`
	Id_rol             uint64 `json:"id_rol"`
	TipoIdentificacion string `json:"tipo_identificacion"`
	Identificacion     string `json:"identificacion" gorm:"size:15;unique;not null"`
	Nombres            string `json:"nombres" gorm:"size:50;"`
	Apellidos          string `json:"apellidos"`
	Correo             string `json:"correo" gorm:"unique;not null"`
	Password           string `json:"password"`
	Fecha_nacimiento   string `json:"fecha_nacimiento"`
	Genero             string `json:"genero"`
	Estado             bool   `json:"estado" gorm:"default:true"`
}

func (u *Usuario) TableName() string {
	return "usuarios"
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerificarPassword(passwordHash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

func (u *Usuario) BeforeSave(tx *gorm.DB) (err error) {
	passwordHash, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(passwordHash)
	return nil
}

func (u *Usuario) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Id_rol = 0
	u.TipoIdentificacion = ""
	u.Identificacion = html.EscapeString(strings.TrimSpace(u.Identificacion))
	u.Nombres = ""
	u.Apellidos = ""
	u.Correo = html.EscapeString(strings.TrimSpace(u.Correo))
	u.Password = ""
	u.Fecha_nacimiento = ""
	u.Genero = ""
	u.Estado = true
	return nil
}

package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	ID                 uint64 `json:"ID" gorm:"primary_key;autoIncrement"`
	Id_rol             uint64 `json:"id_rol"`
	Identificacion     string `json:"identificacion" gorm:"size:15;unique;not null"`
	Nombres            string `json:"nombres" gorm:"size:150;"`
	Apellidos          string `json:"apellidos gorm:"size:150;"`
	Correo             string `json:"correo" gorm:"unique;not null; size:50;"`
	Telefono           string `json:"telefono gorm:"size:50;"`
	Fecha_nacimiento   string `json:"fecha_nacimiento"`
	Genero             string `json:"genero"`
	Password           string `json:"password"`
	Estado             bool   `json:"estado" gorm:"default:true"`
	CitaPaciente       []Cita  `json:"cita_paciente" gorm:"foreignKey:Id_paciente"`
	CitaEspecialista   []Cita  `json:"cita_especialista" gorm:"foreignKey:Id_especialista"`
	gorm.Model
}

func (u *Usuario) TableName() string {
	return "usuarios"
}

func HashPassword(password string) ([]byte, error) {
    if password == "" {
        return nil, nil
    }
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Update the VerificarPassword function to be more robust
func VerificarPassword(hashedPassword, password string) error {
    if hashedPassword == "" || password == "" {
        return bcrypt.ErrMismatchedHashAndPassword
    }
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *Usuario) BeforeSave(tx *gorm.DB) (err error) {
    // Only hash if password is not already hashed (bcrypt hashes are typically 60 chars)
    if len(u.Password) > 0 && len(u.Password) < 60 {
        passwordHash, err := HashPassword(u.Password)
        if err != nil {
            return err
        }
        u.Password = string(passwordHash)
    }
    return nil
}

func (u *Usuario) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Id_rol = 0
	u.Identificacion = html.EscapeString(strings.TrimSpace(u.Identificacion))
	u.Nombres = ""
	u.Apellidos = ""
	u.Correo = html.EscapeString(strings.TrimSpace(u.Correo))
	u.Telefono = ""
	u.Password = ""
	u.Fecha_nacimiento = ""
	u.Genero = ""
	u.Estado = true
	return nil
}

package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type Rol struct {
	gorm.Model

	ID              uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	Id_especialidad uint      `gorm:"not null" json:"id_especialidad"`
	Nombre          string    `gorm:"size:255;not null" json:"nombre"`
	Estado          bool      `gorm:"default:true" json:"estado"`
	Usuarios        []Usuario `json:"usuarios" gorm:"foreignKey:Id_rol"`
}

func (r *Rol) TableName() string {
	return "roles"
}

func (u *Rol) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Nombre = html.EscapeString(strings.TrimSpace(u.Nombre))
	u.Estado = true
	return nil
}

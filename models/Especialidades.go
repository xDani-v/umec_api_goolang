package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type Especialidades struct {

	ID              uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	Id_Rol			uint      `gorm:"size:50;not null" json:"Id_rol_especialidad"`
	Nombre          string    `gorm:"size:50;not null" json:"nombre"`
	Ubicacion 		string    `gorm:"size:150;not null" json:"permiso"`
	Estado          bool      `gorm:"default:true" json:"estado"`
	gorm.Model
}

func (r *Especialidades) TableName() string {
	return "especialidades"
}

func (u *Especialidades) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Nombre = html.EscapeString(strings.TrimSpace(u.Nombre))
	u.Ubicacion = html.EscapeString(strings.TrimSpace(u.Ubicacion))
	u.Estado = true
	return nil
}

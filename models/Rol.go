package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type Rol struct {

	ID              uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	Nombre          string    `gorm:"size:50;not null" json:"nombre"`
	Estado          bool      `gorm:"default:true" json:"estado"`
	Usuarios        []Usuario `json:"usuarios" gorm:"foreignKey:Id_rol"`
    Especialidades  []Especialidades `json:"especialidades" gorm:"foreignKey:Id_Rol"`
	Funcionalidades []RolesFuncionalidad  `json:"rolesfuncionalidad" gorm:"foreignKey:Id_rol"`
	gorm.Model
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

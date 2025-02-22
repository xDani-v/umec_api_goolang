package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type Funcionalidad struct {

	ID              uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	Id_menu         uint      `gorm:"size:50;not null" json:"id_menu"`
	Nombre          string    `gorm:"size:50;not null" json:"nombre"`
	Icono           string    `gorm:"size:0;not null" json:"icono"`
	Ruta            string    `gorm:"size:50;not null" json:"ruta"`
	Estado          bool      `gorm:"default:true" json:"estado"`
	Funcionalidades []RolesFuncionalidad  `json:"rolesfuncionalidad" gorm:"foreignKey:Id_funcionalidad"`
	gorm.Model
}

func (r *Funcionalidad) TableName() string {
	return "funcionalidad"
}

func (u *Funcionalidad) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Nombre = html.EscapeString(strings.TrimSpace(u.Nombre))
	u.Estado = true
	return nil
}

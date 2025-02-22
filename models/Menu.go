package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type Menu struct {

	ID              uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	Nombre          string    `gorm:"size:50;not null" json:"nombre"`
	Icono           string    `gorm:"size:0;not null" json:"icono"`
	Estado          bool      `gorm:"default:true" json:"estado"`
	Funcionalidades  []Funcionalidad `json:"funcionalidades" gorm:"foreignKey:Id_menu"`
	gorm.Model
}

func (r *Menu) TableName() string {
	return "menu"
}

func (u *Menu) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Nombre = html.EscapeString(strings.TrimSpace(u.Nombre))
	u.Estado = true
	return nil
}

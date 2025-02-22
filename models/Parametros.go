package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type Parametros struct {
	ID                uint              `gorm:"primaryKey;autoIncrement" json:"ID"`
	Codigo            string            `gorm:"size:255;not null" json:"codigo"`
	Estado            bool              `gorm:"default:true" json:"estado"`
	ParametrosValores []ParametrosValor `gorm:"foreignKey:IdParametro" json:"parametros_valores"`
	gorm.Model
}

func (r *Parametros) TableName() string {
	return "Parametros"
}

func (u *Parametros) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Codigo = html.EscapeString(strings.TrimSpace(u.Codigo))
	u.Estado = true
	return nil
}

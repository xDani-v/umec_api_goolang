package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type ParametrosValor struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey;autoIncrement" json:"ID"`
	IdParametro uint64 `json:"id_parametro"`
	Valor       string `gorm:"size:255;not null" json:"valor"`
	Estado      bool   `gorm:"default:true" json:"estado"`
}

func (r *ParametrosValor) TableName() string {
	return "ParametrosValor"
}

func (u *ParametrosValor) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Valor = html.EscapeString(strings.TrimSpace(u.Valor))
	u.Estado = true
	return nil
}

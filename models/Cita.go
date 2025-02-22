package models

import (
	 "time"

	"gorm.io/gorm"
)

type Cita struct {

	ID              uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	Id_paciente          string    `gorm:"size:50;not null" json:"id_paciente"`
	Id_especialista 		string    `gorm:"size:150;not null" json:"id_especialista"`
	Fecha_agendamiento    time.Time     `gorm:"size:20;not null" json:"fecha_agendamiento"`
	Hora  				time.Time     `gorm:"size:10;not null" json:"hora"`
	Descripcion  		string    `gorm:"size:300;not null" json:"descripcion"`
	Estado          bool      `gorm:"default:true" json:"estado"`
	gorm.Model
}

func (r *Cita) TableName() string {
	return "cita"
}

func (u *Cita) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Estado = true
	return nil
}

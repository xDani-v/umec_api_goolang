package models

import (

	"gorm.io/gorm"
)

type RolesFuncionalidad struct {

	ID              uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	Id_rol          uint      `gorm:"size:50;not null" json:"id_rol"`
	Id_funcionalidad uint      `gorm:"size:50;not null" json:"id_funcionalidad"`
	gorm.Model
}

func (r *RolesFuncionalidad) TableName() string {
	return "rolesfuncionalidad"
}

func (u *RolesFuncionalidad) Prepare(tx *gorm.DB) (err error) {
	u.ID = 0
	return nil
}

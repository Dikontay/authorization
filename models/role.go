package models

import (
	"auth/pkg"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"unique"`
	Description string `gorm:"size:255;not null" json:"description"`
}

func CreateRole(Role *Role) (err error) {
	err = pkg.DB.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRoles(Roles *[]Role) (err error) {

	return pkg.DB.Find(Roles).Error
}

func GetRole(Role *Role, id int) (err error) {

	return pkg.DB.Where("id = ?", id).First(Role).Error
}

func UpdateRole(Role *Role) {
	_ = pkg.DB.Save(Role)
}

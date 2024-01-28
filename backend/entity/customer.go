package entity

import (
	"gorm.io/gorm"
)
type Ccode struct {
	gorm.Model
	Code string

	Customers	[]Customer `gorm:"foreignKey:CcodeID"`
}

type Cusdetail struct {
	gorm.Model
	Cdname string

	Customers	[]Customer `gorm:"foreignKey:CusdetailID"`
}

type Customer struct {
	gorm.Model

	Name	string

	CcodeID *uint
	Ccode	Ccode	`gorm:"references:id"`

	CusdetailID *uint
	Cusdetail	Cusdetail	`gorm:"references:id"`
}
package entity

import (
	"time"

	"gorm.io/gorm"
)
type Pcode struct {
	gorm.Model
	Name string

	Products	[]Product `gorm:"foreignKey:PcodeID"`
}

type Detail struct {
	gorm.Model
	Name string

	Products	[]Product `gorm:"foreignKey:DetailID"`
}


type Product struct {
	gorm.Model
	Time time.Time

	Name string

	Price int

	Amount int
	
	PcodeID *uint
	Pcode Pcode `gorm:"references:id"`

	DetailID *uint
	Detail Detail `gorm:"references:id"`
}
package models

import "gorm.io/gorm"

type Dna struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Sequence string `gorm:"unique"`
	IsMutant bool   `gorm:"default:false"`
	Size     int
}

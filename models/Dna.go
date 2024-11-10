package models

type Dna struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Sequence string `gorm:"unique"`
	IsMutant bool   `gorm:"default:false"`
	Size     int
}

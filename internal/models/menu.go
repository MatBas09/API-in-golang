package models

type MenuItem struct {
	ID          uint `gorm:"primarykey;autoIncrement"`
	Name        string `gorm:"not null"`
	Type        string 
	Description string
	IsAvailable bool `gorm:"not null"`
	Quantity    int
	Price       float64 `gorm:"not null"`
}
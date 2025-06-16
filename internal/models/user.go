package models

type User struct {
	ID       uint `gorm:"primarykey;autoIncrement"`
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
package model

type User struct {
	ID int64 `gorm:"primaryKey;autoIncrement;not_null"`
	UserName string `gorm:"unique_index;nut_null"`
	FirstName string
	HashPassword string

}

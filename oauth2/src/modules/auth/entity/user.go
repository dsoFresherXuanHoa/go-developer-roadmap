package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"column:email;unique;not null"`
}

type UserResponse struct {
	gorm.Model
	Email string `json:"email" gorm:"column:email;unique;not null"`
}

type UserCreatable struct {
	gorm.Model
	Email *string `json:"email" gorm:"column:email;unique;not null"`
}

type Users []User

func (User) GetTableName() string          { return "users" }
func (UserCreatable) GetTableName() string { return User{}.GetTableName() }
func (Users) GetTableName() string         { return User{}.GetTableName() }

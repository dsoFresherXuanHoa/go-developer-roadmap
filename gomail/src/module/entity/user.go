package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Email      *string `json:"email" gorm:"column:email;unique;not null"`
	Password   string  `json:"-" gorm:"column:password;not null"`
	VerifyCode int     `json:"-" gorm:"column:verify_code;not null"`
}

type UserCreatable struct {
	gorm.Model `json:"-"`
	Email      *string `json:"email" gorm:"column:email;unique;not null"`
	Password   *string `json:"password" gorm:"column:password;not null"`
	VerifyCode int     `json:"-" gorm:"column:verify_code;not null"`
}

type UserUpdatable struct {
	gorm.Model `json:"-"`
	Password   *string `json:"password" gorm:"column:password;not null"`
}

type Users []User

func (User) GetTableName() string          { return "users" }
func (UserCreatable) GetTableName() string { return User{}.GetTableName() }
func (UserUpdatable) GetTableName() string { return User{}.GetTableName() }
func (Users) GetTableName() string         { return User{}.GetTableName() }

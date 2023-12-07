package entity

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	InvalidAccountBadRequest = "Invalid request format: include email, password, avatarURL request!"
	CannotCreateAccount      = "Cannot create account right now: cannot create account, check your request again!"
	CreateAccountSuccess     = "Create account success: congrats!"
)

type Account struct {
	gorm.Model `json:"-"`
	Email      string `json:"email" gorm:"column:email;unique;not null" valid:"email~invalid email format,required~email cannot be empty" `
	Password   string `json:"password" gorm:"column:password;not null" valid:"length(8|16)~password length must be between 8 and 16,required~password cannot be empty"`
	AvatarURL  string `json:"-" form:"avatar" gorm:"column:avatar_url"`
}

type AccountResponse struct {
	Email     string `json:"email"`
	AvatarURL string `json:"avatarURL"`
}

type AccountCreatable struct {
	gorm.Model `json:"-"`
	Email      *string `json:"email" gorm:"column:email;unique;not null" valid:"email~invalid email format,required~email cannot be empty" `
	Password   *string `json:"password" gorm:"column:password;not null" valid:"length(8|16)~password length must be between 8 and 16,required~password cannot be empty"`
	AvatarURL  *string `json:"-" form:"avatar" gorm:"column:avatar_url;"`
}

type AccountUpdatable struct {
	gorm.Model `json:"-"`
	Password   *string `json:"password" gorm:"column:password;not null" valid:"length(8|16)~password length must be between 8 and 16,required~password cannot be empty"`
}

func (account AccountCreatable) Validate() error {
	if _, err := govalidator.ValidateStruct(account); err != nil {
		fmt.Println("Error while validate accountCreatable in repository: " + err.Error())
		return err
	}
	return nil
}

func (account AccountUpdatable) Validate() error {
	if _, err := govalidator.ValidateStruct(account); err != nil {
		fmt.Println("Error while validate accountCreatable in repository: " + err.Error())
		return err
	}
	return nil
}

func (Account) GetTableName() string          { return "accounts" }
func (AccountCreatable) GetTableName() string { return Account{}.GetTableName() }
func (AccountUpdatable) GetTableName() string { return Account{}.GetTableName() }

func (account AccountCreatable) Mask() *AccountCreatable {
	hashPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(*account.Password), 5)
	hashPassword := string(hashPasswordBytes)
	return &AccountCreatable{Email: account.Email, Password: &hashPassword, AvatarURL: account.AvatarURL}
}

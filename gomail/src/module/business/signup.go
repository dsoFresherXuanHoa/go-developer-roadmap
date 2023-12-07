package business

import (
	"context"
	"fmt"
	"go-gomail/src/module/entity"
	"math/rand"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type CreateStorage interface {
	CreateUser(ctx context.Context, user *entity.UserCreatable) (*uint, error)
}

type createBusiness struct {
	storage CreateStorage
}

func NewCreateBusiness(storage CreateStorage) *createBusiness {
	return &createBusiness{storage: storage}
}

func (business *createBusiness) CreateUser(ctx context.Context, user *entity.UserCreatable) (*uint, error) {
	hashPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(*user.Password), 5)
	hashPassword := string(hashPasswordBytes)
	user.Password = &hashPassword
	// TODO: Validate
	verifyCode := GenerateVerifyCode()
	user.VerifyCode = verifyCode
	if err := SendVerifyCode("training.dso.xuanhoa@gmail.com", *user.Email, nil, "Simple Email by Gomail", "Your verify code is: "+strconv.Itoa(verifyCode)); err != nil {
		fmt.Println("Error while create user at send verify code: " + err.Error())
		return nil, err
	} else if id, err := business.storage.CreateUser(ctx, user); err != nil {
		fmt.Println("Error while create user at save db: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func GenerateVerifyCode() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(1000) + 9000
}

func SendVerifyCode(from string, to string, cc *string, sub string, body string) error {
	host := "smtp.gmail.com"
	port := 587
	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", sub)
	m.SetBody("text/plain", body)
	if cc != nil {
		m.SetAddressHeader("Cc", *cc, *cc)
	}

	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error while send mail: " + err.Error())
		return err
	}
	return nil
}

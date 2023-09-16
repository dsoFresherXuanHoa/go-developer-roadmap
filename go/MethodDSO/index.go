package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Username, Password string
}

func (account Account) doAs(username, password string) (*string, error) {
	if username == password {
		return &username, nil
	}
	return nil, errors.New("USERNAME OR PASSWORD IS INCORRECT")
}

func (account *Account) changePassword(password string) string {
	account.Password = password
	return password
}

func main() {
	iset := Account{Username: "dsoFresherXuanHoa", Password: "dsoFresherXuanHoa"}

	if user, err := iset.doAs("root", "root"); err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("Do as: " + *user)
	}

	fmt.Println("ISET: ", iset)
	iset.changePassword("iset")
	fmt.Println("ISET: ", iset)
}

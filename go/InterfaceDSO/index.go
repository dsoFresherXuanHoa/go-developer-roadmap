package main

import (
	"fmt"
	"reflect"
)

type Changeable interface {
	changePassword(password string) string
}

type Account struct {
	Username, Password string
}

type Vehicle struct {
	Password string
}

func (account *Account) changePassword(password string) string {
	account.Password = password
	return password
}

func (vehicle *Vehicle) changePassword(password string) string {
	vehicle.Password = password
	return password
}

func main() {
	iset := Account{Username: "dsoFresherXuanHoa", Password: "dsoFresherXuanHoa"}
	volvo := Vehicle{Password: "dsoFresherXuanHoa"}
	fmt.Println(reflect.TypeOf(iset))
	fmt.Println(reflect.TypeOf(volvo))
}

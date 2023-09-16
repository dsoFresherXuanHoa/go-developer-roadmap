package main

import "fmt"

type Account struct {
	Username, Password string
}

func main() {
	iset := Account{Username: "dsoFresherXuanHoa", Password: "dsoFresherXuanHoa"}
	fmt.Println(iset)
}

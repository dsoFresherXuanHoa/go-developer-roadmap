package main

import (
	"errors"
	"fmt"
)

func whoAmI() {
	fmt.Println("dsoFresherXuanHoa")
}

func doAs(username, password string) (*string, error) {
	if username == password {
		return &username, nil
	}
	return nil, errors.New("USER NAME OR PASSWORD IS INCORRECT")
}

func main() {
	whoAmI()
	username := "dsoFresherXuanHoa"
	password := "dsoFresherXuanHoa"
	if user, err := doAs(username, password); err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("Success: Do as " + *user)
	}
}

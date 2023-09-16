package main

import (
	"errors"
	"fmt"
)

func doAs(username, password string) (*string, error) {
	if username == password {
		return &username, nil
	}
	return nil, errors.New("USERNAME OR PASSWORD IS INCORRECT")
}

func main() {
	username := "dsoFresherXuanHoa"
	password := "dsoFresherXuanHoa"

	if user, err := doAs(username, password); err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("Success: Do as " + *user)
	}

}

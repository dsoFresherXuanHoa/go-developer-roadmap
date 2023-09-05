package main

import "fmt"

func main()  {
	var gender bool = true;
	var author string = "dsoFresherXuanHoa"
	var age int = 22;
	var salary float32 = 500.5;
	var complex complex64 = 1 + 2i;

	fmt.Println(gender, author, age, salary, complex)

	var username string = "Lê Xuân Hòa";
	var usernameRune = []rune(username);
	for _, v := range username {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	for i := 0; i < len(username); i++ {
		fmt.Printf("%c", username[i])
	}
	fmt.Println()
	for i := 0; i < len(usernameRune); i++ {
		fmt.Printf("%c", usernameRune[i])
	}
	fmt.Println()
}
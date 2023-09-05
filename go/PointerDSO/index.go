package main

import (
	"fmt"
	"reflect"
)


func pointerAccess(p *int) {
	*p = 0;
}

func main() {
	var author = "dsoFresherXuanHoa"
	var pAuthor *string = &author

	*pAuthor = "dsoInternXuanHoa"

	fmt.Println(pAuthor)
	fmt.Println(*pAuthor)

	fmt.Println(reflect.TypeOf(pAuthor))

	var age = 22;
	var pAge *int = &age;

	pointerAccess(pAge)

	fmt.Println(age, pAge)
}
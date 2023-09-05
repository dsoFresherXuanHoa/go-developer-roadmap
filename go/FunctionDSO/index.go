package main

import (
	"fmt"
	"time"
)

func whoAmI() {
	fmt.Println("dsoFresherXuanHoa")
}

func doAs(username string, password string) bool {
	return username == password;
}

func timeIt() ( time.Time, string ) {
	return time.Now(), time.Now().GoString()
}

func main() {
	whoAmI();
	fmt.Println(doAs("dsoFresherXuanHoa", "dsoFresherXuanHoa"))
	unix, timeString := timeIt();
	fmt.Println(unix, timeString)	
}
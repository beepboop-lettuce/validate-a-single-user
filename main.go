package main

import (
	"fmt"
	"os"
)

const (
	nothingTyped = "Usage: [username] [password]"
	denied       = "Wrong username or password, try again"
	user1, user2 = "jack", "quentin"
	pswd1, pswd2 = "1888", "1234"
	granted      = "Access granted!"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(nothingTyped)
	} else if os.Args[1] == user1 && os.Args[2] == pswd1 {
		fmt.Println(granted)
	} else if os.Args[1] == user2 && os.Args[2] == pswd2 {
		fmt.Println(granted)
	} else {
		fmt.Println(denied)
	}
}

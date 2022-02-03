package main

import (
	"fmt"
	"os"
)

const (
	nothingTyped = "Usage: [username] [password]"
	deniedUname  = "Access denied, wrong username"
	deniedPswd   = "Access denied, wrong password"
	user1, user2 = "jack", "quentin"
	pswd1, pswd2 = "1888", "1234"
	granted      = "Access granted for %q!\n"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(nothingTyped)
	} else if args[1] == user1 && args[2] == pswd1 {
		fmt.Printf(granted, args[1])
	} else if args[1] == user2 && args[2] == pswd2 {
		fmt.Printf(granted, args[1])
	} else if args[1] != user1 && args[1] != user2 {
		fmt.Println(deniedUname)
	} else {
		fmt.Println(deniedPswd)
	}
}

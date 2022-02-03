package main

import (
	"fmt"
	"os"
)

const (
	nothingTyped = "Usage: [username] [password]"
	wrongUser    = "wrong username"
	wrongPswd    = "wrong password"
	user         = "jack"
	pswd         = "1888"
)

func main() {
	if len(os.Args[1]) < 3 {
		fmt.Println(nothingTyped)
	} else if os.Args[1] != user {
		fmt.Println(wrongUser)
	} else if os.Args[2] != pswd {
		fmt.Println(wrongPswd)
	}
}

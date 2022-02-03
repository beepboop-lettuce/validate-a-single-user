package main

import "os"

const (
	blank := "Usage: [username] [password]",
	wrongUser := "wrong username",
	wrongPswd := "wrong password",
	user := "jack",
	pswd := "1888",

)


func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %q\n", blank)
	}
}

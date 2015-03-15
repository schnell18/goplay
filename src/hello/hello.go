package main

import "fmt"
import "os"
import "strings"

func main() {
	who := "World!"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}

// vim: set ai nu nobk expandtab sw=4 tw=72 syntax=go :

package main

import (
	"fmt"
  // must get package gopass -> `go get github.com/howeyc/gopass`
	"github.com/howeyc/gopass"
)

func main() {
	fmt.Printf("Password: ")
  // ensures password field is hidden while user
  // is typing, returns byte slice
	pass, _ := gopass.GetPasswd()
	fmt.Println(string(pass))
}

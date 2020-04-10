package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	// Perform a Get request.
	resp, err := http.Get("https://raw.githubusercontent.com/binexisHATT/Botnet-Command-Control/master/scripts/net/cc.py")
	// Handling errors.
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // Failure exit.
	}
	// Reading the output from the request and storing it in a variable.
	text, err := ioutil.ReadAll(resp.Body)
	// Handling errors.
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // Failure exit.
	} else {
		// Print the string version of the Get requests body.
		fmt.Println(string(text))
		os.Exit(0)
	}
}
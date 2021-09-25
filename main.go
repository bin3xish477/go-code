package main

import (
	"fmt"

	scriptish "github.com/ganbarodigital/go_scriptish"
)

func main() {
	pipeLine := scriptish.NewPipeline(
		scriptish.Exec("ls", "-la"),
		scriptish.Grep("main"),
	).Exec()
	stdout, _ := pipeLine.TrimmedString()
	fmt.Println(stdout)
}

package main

import (
	"github.com/ian-howell/fooCommand/cmd"
)

func main() {
	rootCommand := cmd.NewFooCommand()
	rootCommand.Execute()
}

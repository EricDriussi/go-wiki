package main

import (
	"wiki/cmd"
	"wiki/src"
)

func main() {
	src.LoadConfig()
	cmd.Execute()
}

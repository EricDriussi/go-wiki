package main

import (
	"wiki/cmd"
	"wiki/pkg"
)

func main() {
	pkg.LoadConfig()
	cmd.Execute()
}

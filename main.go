package main

import (
	"wiki/cmd"
	"wiki/pkg/config"
)

func main() {
	config.Load()
	cmd.Execute()
}

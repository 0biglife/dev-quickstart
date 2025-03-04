package main

import (
	"github.com/0biglife/dev-quickstart/cmd"
	"github.com/0biglife/dev-quickstart/utils"
)

func main() {
	utils.PrintInfo("Starting Dev Quickstart CLI...")
	cmd.Execute()
}

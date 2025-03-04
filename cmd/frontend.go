package cmd

import (
	"os"
	"os/exec"

	"github.com/0biglife/dev-quickstart/utils"
)

func setupFrontend() {
	utils.PrintSetup("Setting up Frontend Project...")

	os.Mkdir("frontend-project", os.ModePerm)

	utils.PrintInfo("Initializing React Project...")
	exec.Command("npx", "create-next-app@latest", "frontend-project").Run()

	utils.PrintInfo("Installing Dependencies...")
	exec.Command("npm", "install", "--save", "redux", "@reduxjs/toolkit", "react-redux", "zustand", "axios", "@apollo/client", "graphql").Run()

	utils.PrintSuccess("Frontend setup completed! Run 'cd frontend-project && npm start' to start.")
}

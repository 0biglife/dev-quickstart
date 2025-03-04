package cmd

import (
	"os"
	"os/exec"

	"github.com/0biglife/dev-quickstart/utils"
)

func setupBackend() {
	utils.PrintSetup("Setting up Backend Project...")

	os.Mkdir("backend-project", os.ModePerm)

	utils.PrintInfo("Initializing NestJS Project...")
	exec.Command("npx", "nest", "new", "backend-project").Run()

	utils.PrintInfo("Installing Dependencies...")
	exec.Command("npm", "install", "--save", "graphql", "@nestjs/graphql", "apollo-server-express", "typeorm", "@nestjs/typeorm", "pg").Run()

	utils.PrintSuccess("Backend setup completed! Run 'cd backend-project && npm run start' to start.")
}

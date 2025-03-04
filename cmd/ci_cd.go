package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/0biglife/dev-quickstart/utils"
)

// CI/CD 설정 함수
func setupCI_CD() {
	utils.PrintSetup("Setting up CI/CD pipeline...")

	// 디렉토리 생성
	os.Mkdir("ci-cd", os.ModePerm)

	// 사용자 입력을 받아 배포 환경 선택
	utils.PrintPrompt("Select deployment platform (github-pages/aws/vercel/netlify/custom): ")
	var deploymentPlatform string
	_, err := fmt.Scanln(&deploymentPlatform)
	if err != nil {
		utils.PrintError("Failed to read input. Defaulting to 'custom'.")
		deploymentPlatform = "custom"
	}

	// CI/CD YAML 파일 생성
	filePath := "ci-cd/github-actions.yaml"
	content := generateCI_CD_YAML(deploymentPlatform)
	utils.WriteFileWithDir(filePath, content)

	utils.PrintSuccess("CI/CD YAML file created at 'ci-cd/github-actions.yaml'.")
}

// 확장성 있는 GitHub Actions YAML 생성 함수
func generateCI_CD_YAML(deploymentPlatform string) string {
	deploymentPlatform = strings.ToLower(deploymentPlatform)
	var deployStep string

	switch deploymentPlatform {
	case "github-pages":
		deployStep = `
      - name: Deploy to GitHub Pages
        run: |
          npm run build
          npx gh-pages -d build
		`
	case "aws":
		deployStep = `
      - name: Deploy to AWS S3
        run: |
          aws s3 sync ./dist s3://my-bucket --delete
		`
	case "vercel":
		deployStep = `
      - name: Deploy to Vercel
        run: vercel --prod
		`
	case "netlify":
		deployStep = `
      - name: Deploy to Netlify
        run: netlify deploy --prod
		`
	default:
		deployStep = `
      - name: Custom Deployment
        run: echo "Define your deployment steps here"
		`
	}

	// GitHub Actions YAML 내용
	yaml := `name: CI/CD Pipeline

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        node-version: [16, 18, 20]
        go-version: [1.18, 1.21]
        python-version: [3.8, 3.10]

    steps:
      - name: Checkout Repository
        uses: actions/checkout@v3

      - name: Setup Node.js
        uses: actions/setup-node@v3
        with:
          node-version: ${{ matrix.node-version }}

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: ${{ matrix.go-version }}

      - name: Setup Python
        uses: actions/setup-python@v4
        with:
          python-version: ${{ matrix.python-version }}

      - name: Install Dependencies
        run: npm install

      - name: Run Linter
        run: npm run lint

      - name: Run Tests
        run: npm test

      - name: Build Project
        run: npm run build

      - name: Cache Dependencies
        uses: actions/cache@v3
        with:
          path: |
            ~/.npm
            ~/.cache/pip
          key: ${{ runner.os }}-deps-${{ hashFiles('**/package-lock.json', '**/go.sum', '**/requirements.txt') }}
          restore-keys: |
            ${{ runner.os }}-deps-

  deploy:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Repository
        uses: actions/checkout@v3

` + deployStep + `

      - name: Notify Deployment Success
        run: echo "🎉 Deployment successful!"`

	return yaml
}

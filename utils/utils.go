package utils

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func PrintSuccess(message string) { fmt.Println(aurora.Green("✅ " + message)) }
func PrintError(message string)   { fmt.Println(aurora.Red("❌ " + message)) }
func PrintInfo(message string)    { fmt.Println(aurora.Cyan("📢 " + message)) }
func PrintPrompt(message string)  { fmt.Print(aurora.Bold("📌 " + message)) }
func PrintSetup(message string)   { fmt.Println(aurora.Magenta("🛠️ " + message)) }

func WriteFileWithDir(filePath, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		PrintError("Failed to create file: " + filePath)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		PrintError("Failed to write to file: " + filePath)
	} else {
		PrintSuccess("File successfully created: " + filePath)
	}
}

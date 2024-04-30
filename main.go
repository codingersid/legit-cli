package main

import (
	"fmt"
	"os"

	"github.com/codingersid/legit-cli/cmd"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: filecreator [filename] [content]")
		os.Exit(1)
	}

	filename := os.Args[1]
	content := os.Args[2]

	err := cmd.CreateFile(filename, content)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("File", filename, "created successfully!")
}

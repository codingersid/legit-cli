package config

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnv() map[string]string {
	env := make(map[string]string)
	file, err := os.Open(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			env[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Error reading .env file")
	}

	return env
}

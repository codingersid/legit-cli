package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func LoadEnv() map[string]string {
	env := make(map[string]string)

	// Path ke file .env, relative terhadap working directory aplikasi
	envFilePath := filepath.Join("..", ".env")

	// Cek apakah file .env ada di lokasi yang ditentukan
	if _, err := os.Stat(envFilePath); os.IsNotExist(err) {
		// Jika file .env tidak ditemukan, kembali ke lokasi root aplikasi
		envFilePath = ".env"
	}

	file, err := os.Open(envFilePath)
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

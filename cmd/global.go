package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
)

// CreateLogFile
func CreateLogFile() *os.File {
	// Path ke file log, relative terhadap working directory aplikasi
	logFilePath := filepath.Join("..", "logs", "logs.log")

	// Buat file log di lokasi yang ditentukan
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Jika file log tidak dapat dibuat di lokasi yang diinginkan, gunakan lokasi dari root aplikasi
		rootPath, err := os.Getwd()
		if err != nil {
			logrus.Fatal(err)
		}
		logFilePath = filepath.Join(rootPath, "logs", "logs.log")
		file, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	return file
}

// CamelToSnake mengonversi string dari CamelCase ke snake_case
func CamelToSnake(s string) string {
	var buf bytes.Buffer
	buf.WriteByte(byte(unicode.ToLower(rune(s[0]))))
	for i := 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			buf.WriteByte('_')
			buf.WriteByte(byte(unicode.ToLower(rune(s[i]))))
		} else {
			buf.WriteByte(byte(s[i]))
		}
	}
	return buf.String()
}

// SnakeCase digunakan untuk mengubah string menjadi format snake_case
func SnakeCase(str string) string {
	var result []rune
	for i, char := range str {
		if unicode.IsUpper(char) {
			// Add underscore before uppercase letters (except the first letter)
			if i > 0 && unicode.IsLower(rune(str[i-1])) {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(char))
		} else if char == ' ' || char == '-' {
			// Replace spaces and hyphens with underscores
			result = append(result, '_')
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

// Mengubah SnakeCase menjadi UC WORD
func SnakeCaseToUCWord(str string) string {
	words := strings.Split(str, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = string(unicode.ToUpper(rune(word[0]))) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// Mengubah SnakeCase menjadi strip
func SnakeCaseToStrip(str string) string {
	return strings.ReplaceAll(str, "_", "-")
}

// UcWordFileFromSnake
func UcWordFileFromSnake(str string) string {
	words := strings.Split(str, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = string(unicode.ToUpper(rune(word[0]))) + word[1:]
		}
	}
	return strings.Join(words, "")
}

// PathToUCWord digunakan untuk mengubah path menjadi format UC Word yang di-merge
func PathToUCWord(path string) string {
	// Split path by slash
	parts := strings.Split(path, "/")
	var result []string
	for _, part := range parts {
		// Convert each part from snake_case to UC Word
		ucWord := UcWordFileFromSnake(part)
		result = append(result, ucWord)
	}
	// Merge all parts together
	return strings.Join(result, "")
}

// TransformPath digunakan untuk views, merubah path terakhir jadi file html
func TransformPath(oldPath string) string {
	// Split the path into components
	parts := strings.Split(oldPath, "/")
	// Ignore the last part of the path
	if len(parts) > 0 {
		parts = parts[:len(parts)-1]
	}
	// Transform each part to lowercase and replace '-' with '_'
	for i, part := range parts {
		parts[i] = SnakeCase(part)
	}
	// Join the parts back into a new path
	newPath := strings.Join(parts, "/")
	return newPath
}

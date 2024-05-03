package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

// ucword
func ucword(s string) string {
	// Menghilangkan karakter non-alphanumeric dan mengonversi ke camel case
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return ""
	}
	words := strings.Split(reg.ReplaceAllString(s, " "), " ")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, "")
}

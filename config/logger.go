package config

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Fungsi untuk menginisialisasi logger
func InitLogger(logFileName string) *logrus.Logger {
	// Path ke file log, relative terhadap working directory aplikasi
	logFilePath := filepath.Join("..", "logs", logFileName+".log")
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

	// Inisialisasi logger
	logger := logrus.New()
	logger.SetOutput(file)
	return logger
}

// Fungsi untuk menutup file log
func CloseLogger(logger *logrus.Logger) {
	file, ok := logger.Out.(*os.File)
	if ok {
		file.Close()
	}
}

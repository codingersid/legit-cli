package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:   "model [nama_model]",
	Short: "Membuat model baru di path: app/models",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		nameParts := strings.Split(modelName, "/")
		modelNameFile := SnakeCase(nameParts[len(nameParts)-1])
		status, err := createModel(modelName)
		if err != nil {
			fmt.Println("Gagal membuat model:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat model, model", modelNameFile, "sudah ada!")
			return
		}
		fmt.Println("Model", modelNameFile, "berhasil dibuat!")
	},
}

func createModel(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	tableName := SnakeCase(nameParts[len(nameParts)-1])
	modelPackageName := PathToUCWord(tableName)

	if strings.Contains(name, "/") {
		return status, errors.New("nama_model yang anda masukkan tidak boleh sebagai path")
	}

	err := os.MkdirAll("app/models", os.ModePerm)
	if err != nil {
		return status, err
	}

	modelPath := fmt.Sprintf("app/models/%s.go", tableName)
	if _, err := os.Stat(modelPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(modelPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	// Isi file model
	code := fmt.Sprintf(`package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// struct ke dan dari database
type %s struct {
    ID        uuid.UUID      `+"`json:\"id\" gorm:\"type:uuid;primaryKey\"`"+`
	// tambahkan kolom tabel lainnya
    CreatedAt time.Time      `+"`json:\"created_at\" gorm:\"default:CURRENT_TIMESTAMP\"`"+`
    UpdatedAt time.Time      `+"`json:\"updated_at\" gorm:\"default:CURRENT_TIMESTAMP\"`"+`
    DeletedAt gorm.DeletedAt `+"`json:\"deleted_at\" gorm:\"index\"`"+`
}

// struct response protection - untuk menampilkan response sesuai kebutuhan saja
type %sResponseProtection struct {
    ID        uuid.UUID      `+"`json:\"id\"`"+`
	// tambahkan kolom tabel lainnya
}

// TableName mengembalikan nama tabel yang digunakan oleh model
func (%s) TableName() string {
    return "%s"
}
`, modelPackageName, modelPackageName, modelPackageName, tableName)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(modelCmd)
}

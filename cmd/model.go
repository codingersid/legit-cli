package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:   "model [nama_model]",
	Short: "Membuat model baru di path: app/models",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		err := createModel(modelName)
		if err != nil {
			fmt.Println("Gagal membuat model:", err)
			return
		}
		fmt.Println("model", modelName, "berhasil dibuat!")
	},
}

func createModel(name string) error {
	nameTable := CamelToSnake(name)
	err := os.MkdirAll("app/models", os.ModePerm)
	if err != nil {
		return err
	}
	modelPath := fmt.Sprintf("app/models/%sModel.go", name)

	file, err := os.Create(modelPath)
	if err != nil {
		return err
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
`, name, name, name, nameTable)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(modelCmd)
}

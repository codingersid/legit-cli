package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration [nama_migration]",
	Short: "Membuat migration baru di path: database/migrations",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrationName := args[0]
		nameParts := strings.Split(migrationName, "/")
		migrationNameFile := SnakeCase(nameParts[len(nameParts)-1])
		status, err := createMigration(migrationName)
		if err != nil {
			fmt.Println("Gagal membuat migration:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat migration, migration", migrationNameFile, "sudah ada!")
			return
		}
		fmt.Println("Migration", migrationNameFile, "berhasil dibuat!")
	},
}

func createMigration(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	tableName := SnakeCase(nameParts[len(nameParts)-1])
	migrationPackageName := PathToUCWord(tableName)

	if strings.Contains(name, "/") {
		return status, errors.New("nama_migration yang anda masukkan tidak boleh sebagai path")
	}

	err := os.MkdirAll("database/migrations", os.ModePerm)
	if err != nil {
		return status, err
	}

	migrationPath := fmt.Sprintf("database/migrations/%s.go", tableName)
	if _, err := os.Stat(migrationPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(migrationPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	// Isi file migration
	code := "package migrations\n\nimport (\n\t\"log\"\n\t\"gorm.io/gorm\"\n)\n\n" +
		"// tabel " + tableName + "\n" +
		"func " + migrationPackageName + "(db *gorm.DB) {\n" +
		"\t// versi 1\n" +
		"\tif !db.Migrator().HasTable(\"" + tableName + "\") {\n" +
		"\t\t// Jika tidak ada, maka buat tabel\n" +
		"\t\tif err := db.Exec(`\n" +
		"\t\t\tCREATE TABLE " + tableName + " (\n" +
		"\t\t\t\tid VARCHAR(36) PRIMARY KEY,\n" +
		"\t\t\t\t// Tambahkan kolom lainnya pada table\n" +
		"\t\t\t\tcreated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,\n" +
		"\t\t\t\tupdated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n" +
		"\t\t\t\tdeleted_at TIMESTAMP\n" +
		"\t\t\t);\n" +
		"\t\t`).Error; err != nil {\n" +
		"\t\t\tlog.Fatal(err)\n" +
		"\t\t}\n" +
		"\t}\n" +
		"}\n"

	_, writeErr := file.WriteString(code)
	if writeErr != nil {
		return status, writeErr
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration [nama_migration]",
	Short: "Membuat migration baru di path: database/migrations",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrationName := args[0]
		err := createMigration(migrationName)
		if err != nil {
			fmt.Println("Gagal membuat migration:", err)
			return
		}
		fmt.Println("Migration", migrationName, "berhasil dibuat!")
	},
}

func createMigration(name string) error {
	nameTable := CamelToSnake(name)
	err := os.MkdirAll("database/migrations", os.ModePerm)
	if err != nil {
		return err
	}
	migrationPath := fmt.Sprintf("database/migrations/%sTable.go", name)

	file, err := os.Create(migrationPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Isi file migration
	code := "package migrations\n\nimport (\n\t\"log\"\n\t\"gorm.io/gorm\"\n)\n\n" +
		"// tabel " + name + "\n" +
		"func " + name + "Table(db *gorm.DB) {\n" +
		"\t// versi 1\n" +
		"\tif !db.Migrator().HasTable(\"" + nameTable + "\") {\n" +
		"\t\t// Jika tidak ada, maka buat tabel\n" +
		"\t\tif err := db.Exec(`\n" +
		"\t\t\tCREATE TABLE " + nameTable + " (\n" +
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
		return writeErr
	}

	return nil
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seeder [nama_seeder]",
	Short: "Membuat seeder baru di path: database/seeders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		seederName := args[0]
		err := createSeeder(seederName)
		if err != nil {
			fmt.Println("Gagal membuat seeder:", err)
			return
		}
		fmt.Println("Seeder", seederName, "berhasil dibuat!")
	},
}

func createSeeder(name string) error {
	nameTable := CamelToSnake(name)
	err := os.MkdirAll("database/seeders", os.ModePerm)
	if err != nil {
		return err
	}
	seederPath := fmt.Sprintf("database/seeders/%sSeeder.go", name)

	file, err := os.Create(seederPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Isi file seeder
	code := fmt.Sprintf(`package seeders

import (
	"log"
	
	//"[ganti_dengan_mod_project_anda]/app/models"
	"gorm.io/gorm"
)

func %sSeeder(db *gorm.DB) {
	// Cek apakah tabel "%s" sudah ada
	if !db.Migrator().HasTable(&models.%s{}) {
		log.Fatal("Tabel '%s' tidak ditemukan.")
	}
}
`, name, nameTable, name, nameTable)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(seederCmd)
}

package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seeder [nama_seeder]",
	Short: "Membuat seeder baru di path: database/seeders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		seederName := args[0]
		nameParts := strings.Split(seederName, "/")
		seederNameFile := SnakeCase(nameParts[len(nameParts)-1])
		status, err := createSeeder(seederName)
		if err != nil {
			fmt.Println("Gagal membuat seeder:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat seeder, seeder", seederNameFile, "sudah ada!")
			return
		}
		fmt.Println("Seeder", seederNameFile, "berhasil dibuat!")
	},
}

func createSeeder(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	tableName := SnakeCase(nameParts[len(nameParts)-1])
	seederPackageName := PathToUCWord(tableName)

	if strings.Contains(name, "/") {
		return status, errors.New("nama_seeder yang anda masukkan tidak boleh sebagai path")
	}

	err := os.MkdirAll("database/seeders", os.ModePerm)
	if err != nil {
		return status, err
	}

	seederPath := fmt.Sprintf("database/seeders/%s.go", tableName)
	if _, err := os.Stat(seederPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(seederPath)
	if err != nil {
		return status, err
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
`, seederPackageName, tableName, tableName, tableName)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(seederCmd)
}

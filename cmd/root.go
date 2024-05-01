// cmd/root.go

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

type MigrationInterface interface {
	RunMigration()
}

var version string

var rootCmd = &cobra.Command{
	Use:   "legit",
	Short: "Legit CLI untuk membuat berbagai dokumen kebutuhan framework legit",
	Long:  `Legit CLI untuk membuat berbagai dokumen kebutuhan framework legit by codingers.id.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Perintah yang anda masukkan salah, silahkan cek legit -h untuk cek bantuan.")
	},
}

var versionCmd = &cobra.Command{
	Use:   "versi",
	Short: "Cek versi legit framework",
	Long:  `Cek versi legit framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Legit Versi :", version)
	},
}

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Menjalankan server dan aplikasi.",
	Run: func(cmd *cobra.Command, args []string) {
		runMain()
	},
}

var runMigrationCmd = &cobra.Command{
	Use:   "runmigration",
	Short: "Menjalankan migration",
	Long:  `Menjalankan migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Panggil fungsi RunMigration() menggunakan path relatif dari direktori saat ini
		// Pastikan untuk menyesuaikan dengan struktur direktori proyek Anda
		err := RunMigration()
		if err != nil {
			fmt.Println("Failed to run migration:", err)
			os.Exit(1)
		}
	},
}

func RunMigration() error {
	// Ganti path relatif dengan path ke file migration.go
	migrationFilePath := "./database/migrations/migration.go"

	// Mengambil direktori saat ini
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Menggabungkan direktori saat ini dengan path file migrasi
	fullPath := filepath.Join(dir, migrationFilePath)

	// Lakukan sesuatu dengan fullPath
	fmt.Println("Running migration from:", fullPath)

	// Implementasi RunMigration() Anda di sini

	return nil
}

func Execute(ver string) {
	version = ver

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(devCmd)
	rootCmd.AddCommand(runMigrationCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runMain() {
	cmd := exec.Command("go", "run", "public/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

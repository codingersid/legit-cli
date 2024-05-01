// cmd/root.go

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

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
		// Mendapatkan path absolut dari file migration.go
		migrationFilePath := filepath.Join("database", "migrations", "migration.go")
		absPath, err := filepath.Abs(migrationFilePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Melakukan dynamic import menggunakan path absolut
		migrationModule, err := os.Open(absPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer migrationModule.Close()

		// Melakukan pemanggilan fungsi RunMigration secara dynamic
		err = findAndRunMigrationFunction(migrationModule)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

// Fungsi untuk mencari dan menjalankan fungsi RunMigration() dari modul migration
func findAndRunMigrationFunction(_ *os.File) error {
	// Implementasi untuk menemukan dan menjalankan fungsi RunMigration()
	migrationType := reflect.TypeOf((*MigrationInterface)(nil)).Elem()
	moduleSymbol := reflect.New(migrationType).Interface().(MigrationInterface)

	moduleValue := reflect.ValueOf(moduleSymbol)

	// Menemukan fungsi RunMigration()
	runMigrationFunc := moduleValue.MethodByName("RunMigration")
	if !runMigrationFunc.IsValid() {
		return fmt.Errorf("RunMigration function not found")
	}

	// Memanggil fungsi RunMigration()
	runMigrationFunc.Call(nil)
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

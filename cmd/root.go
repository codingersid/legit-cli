// cmd/root.go

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type CliRootConfig struct {
	FuncMigration interface{}
	FuncSeeder    interface{}
}

type MigrationInterface interface {
	RunMigration()
}

var version string
var migration interface{}

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

var runMigrationCmd = &cobra.Command{
	Use:   "runmigration",
	Short: "Menjalankan migration",
	Long:  `Menjalankan migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if m, ok := migration.(MigrationInterface); ok {
			m.RunMigration()
		} else {
			fmt.Println("Invalid migration type")
		}
	},
}

func Execute(ver string) {
	version = ver

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runMigrationCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

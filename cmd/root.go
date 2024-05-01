// cmd/root.go

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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

func Execute(ver string) {
	version = ver
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// cmd/root.go

package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Menjalankan server dan aplikasi.",
	Run: func(cmd *cobra.Command, args []string) {
		runMain()
	},
}

func Execute(ver string) {
	version = ver

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(devCmd)

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

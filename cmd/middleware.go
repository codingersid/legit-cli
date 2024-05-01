package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var middlewareCmd = &cobra.Command{
	Use:   "middleware [nama_middleware]",
	Short: "Membuat middleware baru di path: app/http/middlewares",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		middlewareName := args[0]
		err := createMiddleware(middlewareName)
		if err != nil {
			fmt.Println("Gagal membuat middleware:", err)
			return
		}
		fmt.Println("Middleware", middlewareName, "berhasil dibuat!")
	},
}

func createMiddleware(name string) error {
	err := os.MkdirAll("app/http/middlewares", os.ModePerm)
	if err != nil {
		return err
	}
	middlewarePath := fmt.Sprintf("app/http/middlewares/%sMiddleware.go", name)

	file, err := os.Create(middlewarePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Isi file middleware
	code := fmt.Sprintf(`package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func %s(c *fiber.Ctx) error {
	// logic untuk middleware
	return c.Next()
}

`, name)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(middlewareCmd)
}

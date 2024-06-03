package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var middlewareCmd = &cobra.Command{
	Use:   "middleware [nama_middleware]",
	Short: "Membuat middleware baru di path: app/http/middlewares",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		middlewareName := args[0]
		nameParts := strings.Split(middlewareName, "/")
		middlewareNameFile := SnakeCase(nameParts[len(nameParts)-1])
		status, err := createMiddleware(middlewareName)
		if err != nil {
			fmt.Println("Gagal membuat middleware:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat middleware, middleware", middlewareNameFile, "sudah ada!")
			return
		}
		fmt.Println("Middleware", middlewareNameFile, "berhasil dibuat!")
	},
}

func createMiddleware(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	middlewareName := SnakeCase(nameParts[len(nameParts)-1])
	middlewarePackageName := PathToUCWord(middlewareName)

	if strings.Contains(name, "/") {
		return status, errors.New("nama_middleware yang anda masukkan tidak boleh sebagai path")
	}

	err := os.MkdirAll("app/http/middlewares", os.ModePerm)
	if err != nil {
		return status, err
	}

	middlewarePath := fmt.Sprintf("app/http/middlewares/%s.go", middlewareName)
	if _, err := os.Stat(middlewarePath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(middlewarePath)
	if err != nil {
		return status, err
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

`, middlewarePackageName)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(middlewareCmd)
}

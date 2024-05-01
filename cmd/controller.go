package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var controllerCmd = &cobra.Command{
	Use:   "controller [nama_controller]",
	Short: "Membuat controller baru di path: app/http/controllers",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]
		err := createController(controllerName)
		if err != nil {
			fmt.Println("Gagal membuat controller:", err)
			return
		}
		fmt.Println("Controller", controllerName, "berhasil dibuat!")
	},
}

func createController(name string) error {
	nameParts := strings.Split(name, "/")
	packageName := strings.ToLower(nameParts[len(nameParts)-1])
	controllerPath := fmt.Sprintf("app/http/controllers/%s/controller.go", strings.ToLower(name))
	err := os.MkdirAll(fmt.Sprintf("app/http/controllers/%s", strings.ToLower(name)), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(controllerPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Isi file controller
	code := fmt.Sprintf(`package %s

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

// Index adalah handler untuk menampilkan halaman utama.
func Index(c *fiber.Ctx) error {
	return c.SendString("Ini adalah halaman utama.")
}

// Store adalah handler untuk tambah data.
func Store(c *fiber.Ctx) error {
	return c.SendString("Ini adalah tambah data.")
}

// View adalah handler untuk melihat detail data.
func View(c *fiber.Ctx) error {
	return c.SendString("Ini adalah melihat detail data.")
}

// Update adalah handler untuk mengubah data.
func Update(c *fiber.Ctx) error {
	return c.SendString("Ini adalah mengubah data.")
}

// Destroy adalah handler untuk menghapus data.
func Destroy(c *fiber.Ctx) error {
	return c.SendString("Ini adalah menghapus data.")
}
`, packageName)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(controllerCmd)
}

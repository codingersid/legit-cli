package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var withRoute bool

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

		if withRoute {
			err := createRoute(controllerName)
			if err != nil {
				fmt.Println("Gagal membuat route:", err)
				return
			}
			fmt.Println("Route untuk", controllerName, "berhasil dibuat!")
		}
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
	"github.com/gofiber/fiber/v2"
)

// Index adalah handler untuk menampilkan halaman utama.
func Index(c *fiber.Ctx) error {
	return c.SendString("Ini adalah halaman utama.")
}

// Create adalah handler untuk tambah data.
func Create(c *fiber.Ctx) error {
	return c.SendString("Ini adalah tambah data.")
}

// Store adalah handler untuk tambah data.
func Store(c *fiber.Ctx) error {
	return c.SendString("Ini adalah tambah data.")
}

// View adalah handler untuk menampilkan detail data.
func View(c *fiber.Ctx) error {
	return c.SendString("Ini adalah menampilkan detail data.")
}

// Edit adalah handler untuk menampilkan halaman edit.
func Edit(c *fiber.Ctx) error {
	return c.SendString("Ini adalah menampilkan halaman edit.")
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

func createRoute(name string) error {
	nameParts := strings.Split(name, "/")
	// routePackageName := ucword(nameParts[len(nameParts)-1])
	routePackageName := ucword(name)
	controllerPackageName := strings.ToLower(nameParts[len(nameParts)-1])
	controllerPath := fmt.Sprintf("app/http/controllers/%s", strings.ToLower(name))
	err := os.MkdirAll("routes/inners", os.ModePerm)
	if err != nil {
		return err
	}
	routePath := fmt.Sprintf("routes/inners/%sRoute.go", routePackageName)

	file, err := os.Create(routePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Isi file route
	code := fmt.Sprintf(`package inners

import (
	"github.com/codingersid/legit/%s"
	"github.com/gofiber/fiber/v2"
)

func %sRoute(c *fiber.App) {
	// resource route
	prefix := c.Group("/%s")
	prefix.Get("/", %s.Index)
	prefix.Get("/create", %s.Create)
	prefix.Post("/create", %s.Store)
	prefix.Get("/view/:id", %s.View)
	prefix.Get("/edit/:id", %s.Edit)
	prefix.Put("/update/:id", %s.Update)
	prefix.Delete("/delete/:id", %s.Destroy)
}
`, controllerPath, routePackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	controllerCmd.Flags().BoolVar(&withRoute, "with-route", false, "Tambahkan route untuk controller")
	rootCmd.AddCommand(controllerCmd)
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// CONTROLLER:API
var withRouteApi bool
var controllerApiCmd = &cobra.Command{
	Use:   "controller:api [nama_controller]",
	Short: "Membuat controller api baru di path: app/http/controllers/api",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]
		pathName := SnakeCase(controllerName)
		controllerPath := fmt.Sprintf("%s/controller.go", pathName)
		status, err := createControllerApi(controllerName)
		if err != nil {
			fmt.Println("Gagal membuat controller api:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat controller api, controller", controllerPath, "sudah ada!")
			return
		}
		fmt.Println("Controller api", controllerPath, "berhasil dibuat!")

		if withRouteApi {
			err := createRouteApi(controllerName)
			if err != nil {
				fmt.Println("Gagal membuat route api:", err)
				return
			}
			fmt.Println("Route untuk controller api", controllerPath, "berhasil dibuat!")
		}
	},
}

func createControllerApi(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathName := SnakeCase(name)
	packageName := SnakeCase(nameParts[len(nameParts)-1])
	pageName := SnakeCaseToUCWord(packageName)
	controllerPath := fmt.Sprintf("app/http/controllers/api/%s/controller.go", pathName)

	err := os.MkdirAll(fmt.Sprintf("app/http/controllers/api/%s", pathName), os.ModePerm)
	if err != nil {
		return status, err
	}

	if _, err := os.Stat(controllerPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(controllerPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	// Isi file controller
	code := fmt.Sprintf(`package %s

import (
	"github.com/gofiber/fiber/v2"
)

// Index adalah handler untuk menampilkan halaman utama.
func Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Index %s",
	})
}

// Store adalah handler untuk tambah data.
func Store(c *fiber.Ctx) error {
	// Ambil semua data dari form
	formData := make(map[string]string)
	c.Context().Request.PostArgs().VisitAll(func(key, value []byte) {
		formData[string(key)] = string(value)
	})

	// Kembalikan data dalam bentuk JSON
	return c.JSON(fiber.Map{
		"status": fiber.StatusCreated,
		"data":   formData,
		"message": "Store %s",
	})
}

// View adalah handler untuk menampilkan detail data.
func View(c *fiber.Ctx) error {
	// ambil parameter id
	Id := c.Params("id")

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"ID":   Id,
		"message": "View %s",
	})
}

// Update adalah handler untuk mengubah data.
func Update(c *fiber.Ctx) error {
	// ambil parameter id
	Id := c.Params("id")

	// Ambil semua data dari form
	formData := make(map[string]string)
	c.Context().Request.PostArgs().VisitAll(func(key, value []byte) {
		formData[string(key)] = string(value)
	})

	// Kembalikan data dalam bentuk JSON
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"ID":   Id,
		"data":   formData,
		"message": "Update %s",
	})
}

// Destroy adalah handler untuk menghapus data.
func Destroy(c *fiber.Ctx) error {
	// ambil parameter id
	Id := c.Params("id")

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"ID":   Id,
		"message": "Destroy %s",
	})
}

`, packageName, pageName, pageName, pageName, pageName, pageName)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, err
}
func createRouteApi(name string) error {
	nameParts := strings.Split(name, "/")
	pathControllerName := SnakeCase(name)
	controllerPackageName := SnakeCase(nameParts[len(nameParts)-1])
	routePackageName := PathToUCWord(pathControllerName)
	routeFileName := SnakeCase(routePackageName)
	routePrefix := SnakeCaseToStrip(pathControllerName)
	controllerPath := fmt.Sprintf("app/http/controllers/api/%s", pathControllerName)

	err := os.MkdirAll("routes/inners_api", os.ModePerm)
	if err != nil {
		return err
	}

	routePath := fmt.Sprintf("routes/inners_api/%s.go", routeFileName)
	file, err := os.Create(routePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Isi file route
	code := fmt.Sprintf(`package inners_api

import (
	"github.com/codingersid/legit/%s"
	"github.com/gofiber/fiber/v2"
)

func %sRoute(api fiber.Router, c *fiber.App) {
	var router fiber.Router
	if api != nil {
		router = api
	} else {
		router = c
	}
	// resource route
	prefix := router.Group("/%s")
	prefix.Get("/", %s.Index)
	prefix.Post("/create", %s.Store)
	prefix.Get("/view/:id", %s.View)
	prefix.Put("/update/:id", %s.Update)
	prefix.Delete("/delete/:id", %s.Destroy)
}
`, controllerPath, routePackageName, routePrefix, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

// CONTROLLER:WEB
var withRouteWeb bool
var controllerWebCmd = &cobra.Command{
	Use:   "controller:web [nama_controller]",
	Short: "Membuat controller web baru di path: app/http/controllers/web",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]
		pathName := SnakeCase(controllerName)
		controllerPath := fmt.Sprintf("%s/controller.go", pathName)
		status, err := createControllerWeb(controllerName)
		if err != nil {
			fmt.Println("Gagal membuat controller web:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat controller web, controller", controllerPath, "sudah ada!")
			return
		}
		fmt.Println("Controller web", controllerPath, "berhasil dibuat!")

		if withRouteWeb {
			err := createRouteWeb(controllerName)
			if err != nil {
				fmt.Println("Gagal membuat route web:", err)
				return
			}
			fmt.Println("Route untuk controller web", controllerPath, "berhasil dibuat!")
		}
	},
}

func createControllerWeb(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathName := SnakeCase(name)
	packageName := SnakeCase(nameParts[len(nameParts)-1])
	pageName := SnakeCaseToUCWord(packageName)
	controllerPath := fmt.Sprintf("app/http/controllers/web/%s/controller.go", pathName)

	err := os.MkdirAll(fmt.Sprintf("app/http/controllers/web/%s", pathName), os.ModePerm)
	if err != nil {
		return status, err
	}

	if _, err := os.Stat(controllerPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(controllerPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	// Isi file controller
	code := fmt.Sprintf(`package %s

import (
	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/gofiber/fiber/v2"
)

// Index adalah handler untuk menampilkan halaman utama.
func Index(c *fiber.Ctx) error {
	data := fiber.Map{
		"Page": "Index %s",
	}
	return legitConfig.Views("sample", data)(c)
}

// Create adalah handler untuk tambah data.
func Create(c *fiber.Ctx) error {
	data := fiber.Map{
		"Page": "Create %s",
	}
	return legitConfig.Views("sample", data)(c)
}

// Store adalah handler untuk tambah data.
func Store(c *fiber.Ctx) error {
	return c.SendString("Store untuk tambah data dengan method POST.")
}

// View adalah handler untuk menampilkan detail data.
func View(c *fiber.Ctx) error {
	// ambil parameter id
	Id := c.Params("id")

	// parsing data
	data := fiber.Map{
		"Page": "View %s : " + Id,
	}
	return legitConfig.Views("sample", data)(c)
}

// Edit adalah handler untuk menampilkan halaman edit.
func Edit(c *fiber.Ctx) error {
	// ambil parameter id
	Id := c.Params("id")

	// parsing data
	data := fiber.Map{
		"Page": "Edit %s: " + Id,
	}
	return legitConfig.Views("sample", data)(c)
}

// Update adalah handler untuk mengubah data.
func Update(c *fiber.Ctx) error {
	return c.SendString("Update untuk mengubah data dengan method PUT.")
}

// Destroy adalah handler untuk menghapus data.
func Destroy(c *fiber.Ctx) error {
	return c.SendString("Destroy untuk menghapus data dengan method DELETE.")
}

`, packageName, pageName, pageName, pageName, pageName)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, err
}

func createRouteWeb(name string) error {
	nameParts := strings.Split(name, "/")
	pathControllerName := SnakeCase(name)
	controllerPackageName := SnakeCase(nameParts[len(nameParts)-1])
	routePackageName := PathToUCWord(pathControllerName)
	routeFileName := SnakeCase(routePackageName)
	routePrefix := SnakeCaseToStrip(pathControllerName)
	controllerPath := fmt.Sprintf("app/http/controllers/web/%s", pathControllerName)

	err := os.MkdirAll("routes/inners_web", os.ModePerm)
	if err != nil {
		return err
	}

	routePath := fmt.Sprintf("routes/inners_web/%s.go", routeFileName)
	file, err := os.Create(routePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Isi file route
	code := fmt.Sprintf(`package inners_web

import (
	"github.com/codingersid/legit/%s"
	"github.com/gofiber/fiber/v2"
)

func %sRoute(web fiber.Router, c *fiber.App) {
	var router fiber.Router
	if web != nil {
		router = web
	} else {
		router = c
	}
	// resource route
	prefix := router.Group("/%s")
	prefix.Get("/", %s.Index)
	prefix.Get("/create", %s.Create)
	prefix.Post("/create", %s.Store)
	prefix.Get("/view/:id", %s.View)
	prefix.Get("/edit/:id", %s.Edit)
	prefix.Put("/update/:id", %s.Update)
	prefix.Delete("/delete/:id", %s.Destroy)
}
`, controllerPath, routePackageName, routePrefix, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName, controllerPackageName)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	controllerApiCmd.Flags().BoolVar(&withRouteApi, "with:route", false, "Tambahkan route untuk controller api")
	controllerWebCmd.Flags().BoolVar(&withRouteWeb, "with:route", false, "Tambahkan route untuk controller web")
	rootCmd.AddCommand(controllerApiCmd, controllerWebCmd)
}

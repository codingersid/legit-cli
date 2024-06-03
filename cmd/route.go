package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ROUTE:API
var routeApiCmd = &cobra.Command{
	Use:   "route:api [nama_route_api]",
	Short: "Membuat route untuk controller api tertentu yang berada di path: app/http/controllers/api",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]
		pathName := SnakeCase(controllerName)
		controllerPath := fmt.Sprintf("%s/controller.go", pathName)
		status, err := createSingleRouteApi(controllerName)

		if err != nil {
			fmt.Println("Gagal membuat route api:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat route api, route untuk controller api", controllerPath, "sudah ada!")
			return
		}

		fmt.Println("Route untuk controller api", controllerPath, "berhasil dibuat.")
	},
}

func createSingleRouteApi(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathControllerName := SnakeCase(name)
	controllerPackageName := SnakeCase(nameParts[len(nameParts)-1])
	routePackageName := PathToUCWord(pathControllerName)
	routeFileName := SnakeCase(routePackageName)
	routePrefix := SnakeCaseToStrip(pathControllerName)
	controllerPath := fmt.Sprintf("app/http/controllers/api/%s", pathControllerName)
	controllerPathFile := fmt.Sprintf("app/http/controllers/api/%s/controller.go", pathControllerName)

	if _, err := os.Stat(controllerPathFile); os.IsNotExist(err) {
		return status, err
	}

	err := os.MkdirAll("routes/inners_api", os.ModePerm)
	if err != nil {
		return status, err
	}

	routePath := fmt.Sprintf("routes/inners_api/%s.go", routeFileName)
	if _, err := os.Stat(routePath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(routePath)
	if err != nil {
		return status, err
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
		return status, err
	}

	status = "success"
	return status, nil
}

// ROUTE:WEB
var routeWebCmd = &cobra.Command{
	Use:   "route:web [nama_route_web]",
	Short: "Membuat route untuk controller web tertentu yang berada di path: app/http/controllers/web",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]
		pathName := SnakeCase(controllerName)
		controllerPath := fmt.Sprintf("%s/controller.go", pathName)
		status, err := createSingleRouteWeb(controllerName)

		if err != nil {
			fmt.Println("Gagal membuat route web:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat route web, route untuk controller web", controllerPath, "sudah ada!")
			return
		}

		fmt.Println("Route untuk controller web", controllerPath, "berhasil dibuat.")
	},
}

func createSingleRouteWeb(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathControllerName := SnakeCase(name)
	controllerPackageName := SnakeCase(nameParts[len(nameParts)-1])
	routePackageName := PathToUCWord(pathControllerName)
	routeFileName := SnakeCase(routePackageName)
	routePrefix := SnakeCaseToStrip(pathControllerName)
	controllerPath := fmt.Sprintf("app/http/controllers/web/%s", pathControllerName)
	controllerPathFile := fmt.Sprintf("app/http/controllers/web/%s/controller.go", pathControllerName)

	if _, err := os.Stat(controllerPathFile); os.IsNotExist(err) {
		return status, err
	}

	err := os.MkdirAll("routes/inners_web", os.ModePerm)
	if err != nil {
		return status, err
	}

	routePath := fmt.Sprintf("routes/inners_web/%s.go", routeFileName)
	if _, err := os.Stat(routePath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(routePath)
	if err != nil {
		return status, err
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
		return status, err
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(routeApiCmd, routeWebCmd)
}

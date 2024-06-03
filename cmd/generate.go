package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// GENERATE:ROUTE-API
var generateRouteApiCmd = &cobra.Command{
	Use:   "generate:route-api",
	Short: "Generate route otomatis untuk controller api yang berada di path: app/http/controllers/api",
	Run: func(cmd *cobra.Command, args []string) {
		root := "app/http/controllers/api"
		subDirs, err := listSubDirectoriesApi(root)
		if err != nil {
			fmt.Println("Gagal generate route api:", err)
			return
		}

		for _, dir := range subDirs {
			status, err := createRouteFileApi(dir)
			if err != nil {
				continue
			}

			if status == "exist" {
				fmt.Println("Gagal membuat route api, route untuk controller", dir, "sudah ada!")
				continue
			}
			fmt.Println("Route api berhasil dibuat untuk controller:", dir)
		}

		fmt.Println("Proses generate route api selesai dilakukan!")
	},
}

func listSubDirectoriesApi(root string) ([]string, error) {
	var subDirs []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != root {
			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			subDirs = append(subDirs, relativePath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return subDirs, nil
}
func createRouteFileApi(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(strings.ReplaceAll(name, "app/http/controllers/api", ""), "/")
	pathControllerName := name
	controllerPackageName := nameParts[len(nameParts)-1]
	routePackageName := PathToUCWord(pathControllerName)
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

	routePath := fmt.Sprintf("routes/inners_api/%sRoute.go", routePackageName)
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
	return status, err
}

// GENERATE:ROUTE-WEB
var generateRouteWebCmd = &cobra.Command{
	Use:   "generate:route-web",
	Short: "Generate route otomatis untuk controller web yang berada di path: app/http/controllers/web",
	Run: func(cmd *cobra.Command, args []string) {
		root := "app/http/controllers/web"
		subDirs, err := listSubDirectoriesWeb(root)
		if err != nil {
			fmt.Println("Gagal generate route web:", err)
			return
		}

		for _, dir := range subDirs {
			status, err := createRouteFileWeb(dir)
			if err != nil {
				continue
			}

			if status == "exist" {
				fmt.Println("Gagal membuat route web, route untuk controller", dir, "sudah ada!")
				continue
			}
			fmt.Println("Route web berhasil dibuat untuk controller:", dir)
		}

		fmt.Println("Proses generate route web selesai dilakukan!")
	},
}

func listSubDirectoriesWeb(root string) ([]string, error) {
	var subDirs []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != root {
			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			subDirs = append(subDirs, relativePath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return subDirs, nil
}

func createRouteFileWeb(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(strings.ReplaceAll(name, "app/http/controllers/web", ""), "/")
	pathControllerName := name
	controllerPackageName := nameParts[len(nameParts)-1]
	routePackageName := PathToUCWord(pathControllerName)
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

	routePath := fmt.Sprintf("routes/inners_web/%sRoute.go", routePackageName)
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
	rootCmd.AddCommand(generateRouteApiCmd, generateRouteWebCmd)
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// PARTIAL:SCRIPT
var partialScriptCmd = &cobra.Command{
	Use:   "partial:script [nama_script]",
	Short: "Membuat file script di partial baru di path: resources/views/layouts/partials/scripts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		partialName := args[0]
		nameParts := strings.Split(partialName, "/")
		pathName := TransformPath(partialName)
		fileHtml := SnakeCase(nameParts[len(nameParts)-1])
		partialPath := fmt.Sprintf(`scripts/%s/%s.html`, pathName, fileHtml)
		status, err := createPartialScript(partialName)
		if err != nil {
			fmt.Println("Gagal membuat file script di partial:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat partial script, file", partialPath, "sudah ada!")
			return
		}
		fmt.Println("Partial script:", partialPath, "berhasil dibuat!")
	},
}

func createPartialScript(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathName := TransformPath(name)
	fileHtml := SnakeCase(nameParts[len(nameParts)-1])
	partName := SnakeCaseToUCWord(fileHtml)
	partialPath := fmt.Sprintf(`resources/views/layouts/partials/scripts/%s/%s.html`, pathName, fileHtml)
	err := os.MkdirAll(fmt.Sprintf("resources/views/layouts/partials/scripts/%s", pathName), os.ModePerm)

	if err != nil {
		return status, err
	}

	if _, err := os.Stat(partialPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(partialPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	code := fmt.Sprintf(`{{define "script_%s"}}
<!-- Javascript %s Disini -->
{{end}}`, fileHtml, partName)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

// PARTIAL:HEADER
var partialHeaderCmd = &cobra.Command{
	Use:   "partial:header [nama_header]",
	Short: "Membuat file header di partial baru di path: resources/views/layouts/partials/headers",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		partialName := args[0]
		nameParts := strings.Split(partialName, "/")
		pathName := TransformPath(partialName)
		fileHtml := SnakeCase(nameParts[len(nameParts)-1])
		partialPath := fmt.Sprintf(`headers/%s/%s.html`, pathName, fileHtml)
		status, err := createPartialHeader(partialName)
		if err != nil {
			fmt.Println("Gagal membuat file header di partial:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat partial header, file", partialPath, "sudah ada!")
			return
		}
		fmt.Println("Partial header:", partialPath, "berhasil dibuat!")
	},
}

func createPartialHeader(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathName := TransformPath(name)
	fileHtml := SnakeCase(nameParts[len(nameParts)-1])
	partName := SnakeCaseToUCWord(fileHtml)
	partialPath := fmt.Sprintf(`resources/views/layouts/partials/headers/%s/%s.html`, pathName, fileHtml)
	err := os.MkdirAll(fmt.Sprintf("resources/views/layouts/partials/headers/%s", pathName), os.ModePerm)

	if err != nil {
		return status, err
	}

	if _, err := os.Stat(partialPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(partialPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	code := fmt.Sprintf(`{{define "header_%s"}}
<!-- Header %s Disini -->
{{end}}`, fileHtml, partName)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

// PARTIAL:SIDEBAR
var partialSidebarCmd = &cobra.Command{
	Use:   "partial:sidebar [nama_sidebar]",
	Short: "Membuat file sidebar di partial baru di path: resources/views/layouts/partials/sidebars",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		partialName := args[0]
		nameParts := strings.Split(partialName, "/")
		pathName := TransformPath(partialName)
		fileHtml := SnakeCase(nameParts[len(nameParts)-1])
		partialPath := fmt.Sprintf(`sidebars/%s/%s.html`, pathName, fileHtml)
		status, err := createPartialSidebar(partialName)
		if err != nil {
			fmt.Println("Gagal membuat file sidebar di partial:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat partial sidebar, file", partialPath, "sudah ada!")
			return
		}
		fmt.Println("Partial sidebar:", partialPath, "berhasil dibuat!")
	},
}

func createPartialSidebar(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	pathName := TransformPath(name)
	fileHtml := SnakeCase(nameParts[len(nameParts)-1])
	partName := SnakeCaseToUCWord(fileHtml)
	partialPath := fmt.Sprintf(`resources/views/layouts/partials/sidebars/%s/%s.html`, pathName, fileHtml)
	err := os.MkdirAll(fmt.Sprintf("resources/views/layouts/partials/sidebars/%s", pathName), os.ModePerm)

	if err != nil {
		return status, err
	}

	if _, err := os.Stat(partialPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(partialPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	code := fmt.Sprintf(`{{define "sidebar_%s"}}
<!-- Sidebar %s Disini -->
{{end}}`, fileHtml, partName)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(partialScriptCmd, partialHeaderCmd, partialSidebarCmd)
}

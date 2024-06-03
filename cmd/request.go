package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var requestCmd = &cobra.Command{
	Use:   "request [nama_request]",
	Short: "Membuat request baru di path: app/http/requests",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		requestName := args[0]
		nameParts := strings.Split(requestName, "/")
		requestNameFile := SnakeCase(nameParts[len(nameParts)-1])
		status, err := createRequest(requestName)
		if err != nil {
			fmt.Println("Gagal membuat request:", err)
			return
		}

		if status == "exist" {
			fmt.Println("Gagal membuat request, request", requestNameFile, "sudah ada!")
			return
		}
		fmt.Println("Request", requestNameFile, "berhasil dibuat!")
	},
}

func createRequest(name string) (string, error) {
	status := "failed"
	nameParts := strings.Split(name, "/")
	requestNameFile := SnakeCase(nameParts[len(nameParts)-1])

	if strings.Contains(name, "/") {
		return status, errors.New("nama_request yang anda masukkan tidak boleh sebagai path")
	}

	err := os.MkdirAll("app/http/requests", os.ModePerm)
	if err != nil {
		return status, err
	}

	requestPath := fmt.Sprintf("app/http/requests/%s.go", requestNameFile)
	if _, err := os.Stat(requestPath); err == nil {
		status = "exist"
		return status, nil
	} else if !os.IsNotExist(err) {
		return status, err
	}

	file, err := os.Create(requestPath)
	if err != nil {
		return status, err
	}
	defer file.Close()

	// Isi file request
	code := fmt.Sprintf(`package %s

// Store Request
type (
	StoreRequest struct {
		ID   string `+"`json:\"id\" validate:\"required,len=2\"`"+`
		// tambahkan struct lainnya
	}
)

// Update Request
type (
	UpdateRequest struct {
		ID   string `+"`json:\"id\" validate:\"required,len=2\"`"+`
		// tambahkan struct lainnya
	}
)
`, requestNameFile)

	_, err = file.WriteString(code)
	if err != nil {
		return status, err
	}

	status = "success"
	return status, nil
}

func init() {
	rootCmd.AddCommand(requestCmd)
}

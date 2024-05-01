package cmd

import (
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
		err := createRequest(requestName)
		if err != nil {
			fmt.Println("Gagal membuat request:", err)
			return
		}
		fmt.Println("Request", requestName, "berhasil dibuat!")
	},
}

func createRequest(name string) error {
	nameParts := strings.Split(name, "/")
	packageName := strings.ToLower(nameParts[len(nameParts)-1])
	requestPath := fmt.Sprintf("app/http/requests/%s/request.go", strings.ToLower(name))
	err := os.MkdirAll(fmt.Sprintf("app/http/requests/%s", strings.ToLower(name)), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(requestPath)
	if err != nil {
		return err
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
`, packageName)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(requestCmd)
}

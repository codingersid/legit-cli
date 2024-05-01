package contoh

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "legit",
	Short: "Legit CLI untuk membuat berbagai dokumen kebutuhan framework legit",
	Long:  `Legit CLI untuk membuat berbagai dokumen kebutuhan framework legit by codingers.id.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gunakan salah satu perintah berikut: controller, model, atau request")
	},
}

var controllerCmd = &cobra.Command{
	Use:   "controller [nama_controllernya]",
	Short: "Membuat controller baru",
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

var modelCmd = &cobra.Command{
	Use:   "model [nama_modelnya]",
	Short: "Membuat model baru",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		err := createModel(modelName)
		if err != nil {
			fmt.Println("Gagal membuat model:", err)
			return
		}
		fmt.Println("model", modelName, "berhasil dibuat!")
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

// Index adalah handler untuk menampilkan halaman utama %s.
func Index(ctx *fiber.Ctx) error {
	return ctx.SendString("Ini adalah halaman utama %s")
}
`, packageName, name, name)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func createModel(name string) error {
	nameTable := camelToSnake(name)
	modelPath := fmt.Sprintf("app/models/%sModel.go", name)

	file, err := os.Create(modelPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Isi file model
	code := fmt.Sprintf(`package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// struct ke dan dari database
type %s struct {
    ID        uuid.UUID      `+"`json:\"id\" gorm:\"type:uuid;primaryKey\"`"+`
	// masukkan kolom lainnya
    CreatedAt time.Time      `+"`json:\"created_at\" gorm:\"default:CURRENT_TIMESTAMP\"`"+`
    UpdatedAt time.Time      `+"`json:\"updated_at\" gorm:\"default:CURRENT_TIMESTAMP\"`"+`
    DeletedAt gorm.DeletedAt `+"`json:\"deleted_at\" gorm:\"index\"`"+`
}

// struct response protection - untuk menampilkan response sesuai kebutuhan saja
type %sResponseProtection struct {
    ID        uuid.UUID      `+"`json:\"id\"`"+`
	// masukkan kolom lainnya
}

// TableName mengembalikan nama tabel yang digunakan oleh model User
func (%s) TableName() string {
    return "%s"
}
`, name, name, name, nameTable)

	_, err = file.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

// camelToSnake mengonversi string dari CamelCase ke snake_case
func camelToSnake(s string) string {
	var buf bytes.Buffer
	buf.WriteByte(byte(unicode.ToLower(rune(s[0]))))
	for i := 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			buf.WriteByte('_')
			buf.WriteByte(byte(unicode.ToLower(rune(s[i]))))
		} else {
			buf.WriteByte(byte(s[i]))
		}
	}
	return buf.String()
}

func LegitCommand() {
	rootCmd.AddCommand(controllerCmd)
	rootCmd.AddCommand(modelCmd)
	// rootCmd.AddCommand(requestCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Hello() {
	fmt.Println("Hello from example package!")
}

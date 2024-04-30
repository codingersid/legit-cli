package cmd

import (
	"os"
	"strings"
)

// CreateFile akan membuat sebuah file dengan nama yang diberikan dan isi teks yang diberikan.
func CreateFile(filename, content string) error {
	// Mencari indeks terakhir dari tanda titik (.)
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex == -1 {
		// Jika tidak ada tanda titik (.) maka tambahkan ".go" sebagai default
		filename += ".go"
	} else if lastDotIndex == 0 || lastDotIndex == len(filename)-1 {
		// Jika tanda titik (.) berada di awal atau akhir string, maka tambahkan ".go" sebagai default
		filename += "go"
	}

	return os.WriteFile(filename, []byte(content), 0644)
}

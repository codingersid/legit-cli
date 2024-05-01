package cmd

import (
	"bytes"
	"unicode"
)

// CamelToSnake mengonversi string dari CamelCase ke snake_case
func CamelToSnake(s string) string {
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

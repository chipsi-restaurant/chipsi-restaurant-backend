package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	result := ""
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			result += "_"
		}
		result += string(unicode.ToLower(r))
	}
	return result
}

func GeneratePromoCode(length int) string {
	b := make([]byte, length/2)
	_, _ = rand.Read(b)
	return strings.ToUpper(hex.EncodeToString(b)) // например: 8E9F2A7B
}

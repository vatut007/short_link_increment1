package utils

import (
	"encoding/base64"
	"fmt"
)

func GenerateShortCodeUrl(originalURL string) string {
	hash := fmt.Sprintf("%x", []byte(originalURL))
	encoded := base64.StdEncoding.EncodeToString([]byte(hash))
	return encoded[:8]
}

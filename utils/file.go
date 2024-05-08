package utils

import (
	"os"
)

// ReadFile reads file from the given path
func ReadFile(filePath string) (body []byte) {

	body,err := os.ReadFile(filePath)

	if err != nil {
		return nil
	}
	return body
}

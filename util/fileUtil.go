package util

import (
	"net/http"
	"os"
)

func SendHtml(writer http.ResponseWriter, path string) error {
	content, err := os.ReadFile(path)
	if (err != nil) {
		return err
	}
	writer.Write(content)

	return nil
}

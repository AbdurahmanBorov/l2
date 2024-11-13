package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// downloadFile загружает содержимое страницы и сохраняет его в файл
func downloadFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Извлекаем имя файла из URL
	fileName := "index.html"
	if strings.HasSuffix(url, "/") {
		fileName = "index.html"
	} else {
		parts := strings.Split(url, "/")
		fileName = parts[len(parts)-1]
		if !strings.Contains(fileName, ".") {
			fileName += ".html"
		}
	}

	// Создаем файл для записи содержимого
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	url := "http://example.com"
	fmt.Println("Downloading:", url)
	if err := downloadFile(url); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Downloaded successfully.")
	}
}

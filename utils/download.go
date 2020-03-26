package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Download URL を渡してダウンロードする
func Download(url string, time time.Time, savePath string, filename string, ext string) (path string, err error) {
	var (
		res  *http.Response
		file *os.File
	)
	path = fmt.Sprintf("%s/%s.%s", savePath, filename, ext)
	if res, err = http.Get(url); err != nil {
		return "", err
	}
	defer res.Body.Close()
	if err = os.MkdirAll(savePath, 0755); err != nil {
		return "", err
	}
	if file, err = os.Create(path); err != nil {
		return "", err
	}
	defer file.Close()
	io.Copy(file, res.Body)
	if err = os.Chtimes(path, time, time); err != nil {
		return "", err
	}
	return path, nil
}

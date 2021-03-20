package fwrite

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

func WriteJson(path string, i interface{}) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	fh, err := os.OpenFile(path, os.O_CREATE|syscall.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer fh.Close()

	return json.NewEncoder(fh).Encode(i)
}

func WriteRaw(filePath string, raw []byte) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, raw, 0664)
}

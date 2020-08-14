package fwrite

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	UseGzip Option = 1
)

type Option uint8

func WriteHandler(saveFile string, handler func(io.Writer) error, options ...Option) error {
	if err := os.MkdirAll(filepath.Dir(saveFile), 0777); err != nil {
		return err
	}

	file, err := os.OpenFile(saveFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return err
	}
	defer file.Close()

	var writer io.WriteCloser
	writer = file

	for _, option := range options {
		switch option {
		case UseGzip:
			if strings.HasSuffix(saveFile, `.gz`) {
				writer, err = gzip.NewWriterLevel(file, gzip.BestCompression)
				if err != nil {
					return err
				}
			}
		}
	}

	defer writer.Close()

	if err = handler(writer); err != nil {
		return err
	}

	return nil
}

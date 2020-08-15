package fwrite

import (
	"io"
)

func WriteString(saveFile string, s string, options ...Option) error {
	return WriteHandler(saveFile, func(writer io.Writer) error {
		_, err := writer.Write([]byte(s))
		return err
	}, options...)
}

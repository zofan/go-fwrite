package fwrite

import (
	"io"
)

func WriteReader(saveFile string, reader io.Reader, options ...Option) error {
	return WriteHandler(saveFile, func(writer io.Writer) error {
		_, err := io.Copy(writer, reader)
		return err
	}, options...)
}

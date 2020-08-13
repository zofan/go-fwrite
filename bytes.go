package fwrite

import (
	"io"
)

func WriteBytes(saveFile string, raw []byte, options ...Option) error {
	return WriteHandler(saveFile, func(writer io.Writer) error {
		_, err := writer.Write(raw)
		return err
	}, options...)
}

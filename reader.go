package fwrite

import (
	"github.com/zofan/go-bits"
	"io"
)

func FromReader(saveFile string, reader io.Reader, options bits.Bits8) error {
	return FromHandler(saveFile, func(writer io.Writer) error {
		_, err := io.Copy(writer, reader)
		return err
	}, options)
}

package fwrite

import (
	"github.com/zofan/go-bits"
	"io"
)

func FromString(saveFile string, s string, options bits.Bits8) error {
	return FromHandler(saveFile, func(writer io.Writer) error {
		_, err := writer.Write([]byte(s))
		return err
	}, options)
}

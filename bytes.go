package fwrite

import (
	"github.com/zofan/go-bits"
	"io"
)

func FromBytes(saveFile string, raw []byte, options bits.Bits8) error {
	return FromHandler(saveFile, func(writer io.Writer) error {
		_, err := writer.Write(raw)
		return err
	}, options)
}

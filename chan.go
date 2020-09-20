package fwrite

import (
	"github.com/zofan/go-bits"
	"io"
)

func FromChan(saveFile string, ch chan []byte, options bits.Bits8) error {
	return FromHandler(saveFile, func(writer io.Writer) error {
		for v := range ch {
			_, err := writer.Write(v)
			if err != nil {
				return err
			}
		}

		return nil
	}, options)
}

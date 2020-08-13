package fwrite

import (
	"io"
)

func WriteChan(saveFile string, ch chan []byte, options ...Option) error {
	return WriteHandler(saveFile, func(writer io.Writer) error {
		for v := range ch {
			_, err := writer.Write(v)
			if err != nil {
				return err
			}
		}

		return nil
	}, options...)
}

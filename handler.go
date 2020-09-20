package fwrite

import (
	"compress/gzip"
	"github.com/zofan/go-bits"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	None    bits.Bits8 = 0
	UseGzip bits.Bits8 = 1
	Append  bits.Bits8 = 2
)

func FromHandler(saveFile string, handler func(io.Writer) error, options bits.Bits8) error {
	if err := os.MkdirAll(filepath.Dir(saveFile), 0777); err != nil {
		return err
	}

	flag := os.O_WRONLY | os.O_CREATE
	if options.Has(Append) {
		flag |= os.O_APPEND
	}

	file, err := os.OpenFile(saveFile, flag, 0664)
	if err != nil {
		return err
	}
	defer file.Close()

	var writer io.WriteCloser
	writer = file

	if options.Has(UseGzip) && strings.HasSuffix(saveFile, `.gz`) {
		writer, err = gzip.NewWriterLevel(file, gzip.BestCompression)
		if err != nil {
			return err
		}
	}

	if writer != file {
		defer writer.Close()
	}

	return handler(writer)
}

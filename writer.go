package fwrite

import (
	"compress/gzip"
	"encoding/json"
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

type Writer struct {
	fh io.WriteCloser
	gz io.WriteCloser
}

func NewWriter(filePath string, options bits.Bits8) (writer *Writer, err error) {
	if err := os.MkdirAll(filepath.Dir(filePath), 0777); err != nil {
		return nil, err
	}

	flag := os.O_WRONLY | os.O_CREATE
	if options.Has(Append) {
		flag |= os.O_APPEND
	}

	writer = &Writer{}

	writer.fh, err = os.OpenFile(filePath, flag, 0664)
	if err != nil {
		return nil, err
	}

	if options.Has(UseGzip) && strings.HasSuffix(filePath, `.gz`) {
		writer.gz, err = gzip.NewWriterLevel(writer.fh, gzip.BestCompression)
		if err != nil {
			return nil, err
		}
	}

	return writer, nil
}

func (w *Writer) Writer() io.Writer {
	if w.gz != nil {
		return w.gz
	}

	return w.fh
}

func (w *Writer) Close() error {
	if w.gz != nil && w.gz != w.fh {
		err := w.gz.Close()
		if err != nil {
			return err
		}
	}

	return w.fh.Close()
}

func (w *Writer) WriteBytes(b []byte) error {
	_, err := w.Writer().Write(b)
	return err
}

func (w *Writer) WriteString(s string) error {
	_, err := w.Writer().Write([]byte(s))
	return err
}

func (w *Writer) WriteJson(v interface{}) error {
	return json.NewEncoder(w.Writer()).Encode(v)
}

func (w *Writer) WriteReader(r io.Reader) error {
	_, err := io.Copy(w.Writer(), r)
	return err
}

func (w *Writer) WriteChan(ch chan []byte) error {
	for v := range ch {
		_, err := w.Writer().Write(v)
		if err != nil {
			return err
		}
	}

	return nil
}

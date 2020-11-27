package fwrite

import (
	"github.com/pkg/errors"
	"github.com/zofan/go-bits"
	"sort"
	"strings"
)

type StringBuffer struct {
	buffer []string
	writer *Writer

	BufferSize int
	BufferSort bool
}

func NewStringBuffer(filePath string, options bits.Bits8) (*StringBuffer, error) {
	options.Set(Append)

	w, err := NewWriter(filePath, options)
	if err != nil {
		return nil, err
	}

	return &StringBuffer{writer: w, BufferSize: 1000}, nil
}

func (b *StringBuffer) Push(v ...string) {
	b.PushSlice(v)
}

func (b *StringBuffer) PushSlice(v []string) {
	if len(b.buffer) >= b.BufferSize {
		b.Flush()
	}

	b.buffer = append(b.buffer, v...)
}

func (b *StringBuffer) Flush() {
	b.buffer = uniqueStrings(b.buffer)

	if b.BufferSort {
		sort.Strings(b.buffer)
	}

	err := b.writer.WriteString(strings.Join(b.buffer, "\n") + "\n")
	if err == nil {
		b.buffer = []string{}
	}
}

func (b *StringBuffer) Close() error {
	b.Flush()

	if len(b.buffer) > 0 {
		return errors.New(`buffer is not empty`)
	}

	return b.writer.Close()
}

func uniqueStrings(s []string) []string {
	uniq := map[string]struct{}{}

	for _, v := range s {
		uniq[v] = struct{}{}
	}

	var result []string
	for key := range uniq {
		result = append(result, key)
	}

	return result
}

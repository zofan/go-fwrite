package fwrite

import (
	"github.com/zofan/go-bits"
	"strings"
)

type StringBuffer struct {
	buffer []string

	bufferSize int
	writer     *Writer
}

func NewStringBuffer(filePath string, bufferSize int, options bits.Bits8) (*StringBuffer, error) {
	options.Set(Append)

	w, err := NewWriter(filePath, options)
	if err != nil {
		return nil, err
	}

	return &StringBuffer{writer: w, bufferSize: bufferSize}, nil
}

func (b *StringBuffer) Push(r string) {
	if len(b.buffer) >= b.bufferSize {
		b.Save()
	}

	b.buffer = append(b.buffer, r)
}

func (b *StringBuffer) PushSlice(r []string) {
	if len(b.buffer) >= b.bufferSize {
		b.Save()
	}

	b.buffer = append(b.buffer, r...)
}

func (b *StringBuffer) Save() {
	b.buffer = uniqueStrings(b.buffer)

	err := b.writer.WriteString(strings.Join(b.buffer, "\n") + "\n")
	if err == nil {
		b.buffer = []string{}
	}
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

package fwrite

import (
	"io"
	"strings"
)

type StringBuffer struct {
	buffer []string

	BufferSize int
	SaveFile   string
}

func (b *StringBuffer) Push(r string) {
	if len(b.buffer) >= b.BufferSize {
		b.Save()
	}

	b.buffer = append(b.buffer, r)
}

func (b *StringBuffer) PushSlice(r []string) {
	if len(b.buffer) >= b.BufferSize {
		b.Save()
	}

	b.buffer = append(b.buffer, r...)
}

func (b *StringBuffer) Save() {
	b.buffer = uniqueStrings(b.buffer)

	err := FromHandler(b.SaveFile, func(writer io.Writer) error {
		_, err := writer.Write([]byte(strings.Join(b.buffer, "\n") + "\n"))
		return err
	}, Append)
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

package fwrite

import (
	"github.com/zofan/go-util"
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
	} else {
		b.buffer = append(b.buffer, r)
	}
}

func (b *StringBuffer) PushSlice(r []string) {
	if len(b.buffer) >= b.BufferSize {
		b.Save()
	} else {
		b.buffer = append(b.buffer, r...)
	}
}

func (b *StringBuffer) Save() {
	b.buffer = util.UniqueStrings(b.buffer)

	err := WriteHandler(b.SaveFile, func(writer io.Writer) error {
		_, err := writer.Write([]byte(strings.Join(b.buffer, "\n") + "\n"))
		return err
	})
	if err == nil {
		b.buffer = []string{}
	}
}

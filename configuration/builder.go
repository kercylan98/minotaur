package configuration

import (
	"fmt"
	"strings"
)

func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct {
	strings.Builder
	err error
}

func (b *Builder) Fprintf(format string, a ...interface{}) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = fmt.Fprintf(&b.Builder, format, a...)
	return b
}

func (b *Builder) WriteRune(r rune) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.Builder.WriteRune(r)
	return b
}

func (b *Builder) WriteBytes(c ...byte) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.Builder.Write(c)
	return b
}

func (b *Builder) WriteString(s string) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.Builder.WriteString(s)
	return b
}

func (b *Builder) String() string {
	return b.Builder.String()
}

func (b *Builder) Reset() {
	b.Builder.Reset()
	b.err = nil
}

func (b *Builder) Error() error {
	return b.err
}

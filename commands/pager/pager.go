package pager

import (
	"io"
	"os"
)

type Pager struct {
	writer io.WriteCloser
}

func New() *Pager {
	return &Pager{
		writer: nopCloser{os.Stdout},
	}
}

func (p *Pager) Writer() io.WriteCloser {
	return p.writer
}

func (p *Pager) Close() error {
	return nil
}

type nopCloser struct{ io.Writer }

func (nopCloser) Close() error { return nil }

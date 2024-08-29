package gzip

import (
	"compress/gzip"
	"io"
)

type compressReader struct {
	r  io.Reader
	zr *gzip.Reader
}

func NewReader(r io.Reader) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

func (c compressReader) Read(b []byte) (int, error) {
	return c.zr.Read(b)
}

func (c *compressReader) Close() error {
	if err := c.zr.Close(); err != nil {
		return err
	}
	return nil
}

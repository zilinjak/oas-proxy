package internal

import (
	"bytes"
	"io"
)

func CopyData(original io.ReadCloser) (io.ReadCloser, io.ReadCloser, error) {
	data, err := io.ReadAll(original)
	if err != nil {
		return nil, nil, err
	}
	copy1 := io.NopCloser(bytes.NewBuffer(data))
	copy2 := io.NopCloser(bytes.NewBuffer(data))

	return copy1, copy2, nil
}

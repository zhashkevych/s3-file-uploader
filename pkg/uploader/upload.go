package uploader

import (
	"context"
	"io"
)

type Uploader interface {
	Upload(ctx context.Context, file io.Reader, size int64, contentType string) (string, error)
}

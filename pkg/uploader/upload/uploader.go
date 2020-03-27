package upload

import (
	"context"
	"github.com/zhashkevych/s3-file-uploader/sidecar/filestorage"
	"io"
	"math/rand"
)

const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fileNameLength = 16
)

type Uploader struct {
	fs *filestorage.FileStorage
}

func NewUploader(fs *filestorage.FileStorage) *Uploader {
	return &Uploader{
		fs: fs,
	}
}

func (u *Uploader) Upload(ctx context.Context, file io.Reader, size int64, contentType string) (string, error) {
	return u.fs.Upload(ctx, filestorage.UploadInput{
		File:        file,
		Name:        generateFileName(),
		Size:        size,
		ContentType: contentType,
	})
}

func generateFileName() string {
	b := make([]byte, fileNameLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

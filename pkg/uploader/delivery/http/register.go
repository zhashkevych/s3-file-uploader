package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/s3-file-uploader/pkg/uploader"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uploader uploader.Uploader) {
	h := NewHandler(uploader)

	router.POST("/upload", h.Upload)
}

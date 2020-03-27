package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/s3-file-uploader/pkg/uploader"
	"net/http"
)

const (
	MAX_UPLOAD_SIZE = 5 << 20 // 5 megabytes
)

var (
	IMAGE_TYPES = map[string]interface{}{
		"image/jpeg": nil,
		"image/png":  nil,
	}
)

type Handler struct {
	uploader uploader.Uploader
}

func NewHandler(uploader uploader.Uploader) *Handler {
	return &Handler{
		uploader: uploader,
	}
}

type uploadResponse struct {
	Status string `json:"status"`
	Msg    string `json:"message,omitempty"`
	URL    string `json:"url,omitempty"`
}

func (h *Handler) Upload(c *gin.Context) {
	// Limit Upload File Size
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MAX_UPLOAD_SIZE)

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &uploadResponse{
			Status: "error",
			Msg:    err.Error(),
		})
		return
	}
	
	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	file.Read(buffer)
	fileType := http.DetectContentType(buffer)

	// Validate File Type
	if _, ex := IMAGE_TYPES[fileType]; !ex {
		c.JSON(http.StatusBadRequest, &uploadResponse{
			Status: "error",
			Msg:    "file type is not supported",
		})
		return
	}

	url, err := h.uploader.Upload(c.Request.Context(), file, fileHeader.Size, fileType)
	if err != nil {
		c.JSON(http.StatusBadRequest, &uploadResponse{
			Status: "error",
			Msg:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &uploadResponse{
		Status: "ok",
		URL:    url,
	})
}

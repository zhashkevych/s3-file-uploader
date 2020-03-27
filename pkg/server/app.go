package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/spf13/viper"

	"github.com/zhashkevych/s3-file-uploader/pkg/uploader"
	"github.com/zhashkevych/s3-file-uploader/pkg/uploader/upload"
	uphttp "github.com/zhashkevych/s3-file-uploader/pkg/uploader/delivery/http"
	"github.com/zhashkevych/s3-file-uploader/sidecar/filestorage"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	fileStorage   *filestorage.FileStorage
	imageUploader uploader.Uploader
}

func NewApp(accessKey, secretKey string) *App {
	// Initiate an S3 compatible client
	client, err := minio.New(viper.GetString("storage.endpoint"), accessKey, secretKey, false)
	if err != nil {
		log.Fatal(err)
	}

	fileStorage := filestorage.NewFileStorage(
		client,
		viper.GetString("storage.bucket"),
		viper.GetString("storage.endpoint"),
	)

	return &App{
		fileStorage:   fileStorage,
		imageUploader: upload.NewUploader(fileStorage),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// API endpoints
	api := router.Group("/api")
	uphttp.RegisterHTTPEndpoints(api, a.imageUploader)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
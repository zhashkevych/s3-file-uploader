package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zhashkevych/s3-file-uploader/pkg/config"
	"github.com/zhashkevych/s3-file-uploader/pkg/server"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	app := server.NewApp(accessKey, secretKey)

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

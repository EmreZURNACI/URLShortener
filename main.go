package main

import (
	"fmt"
	"log"
	"time"

	"github.com/EmreZURNACI/url-shortener/infra"
	"github.com/EmreZURNACI/url-shortener/server"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./.config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	zap.L().Info("application starting...")

}

func main() {

	handler, err := infra.Connection()

	if err != nil {
		log.Fatal(err.Error())
	}
	if err := handler.CreateTable(); err != nil {
		log.Fatal(err.Error())
	}

	server.Start(handler)

}

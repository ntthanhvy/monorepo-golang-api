package main

import (
	"time"
	"os"
	"github.com/go-kit/kit/log"
	"core-api/handler"
	"net/http"

	"github.com/joho/godotenv"
)

var (
  httpAddr = ":8080"
)

func main() {
  logger := log.NewJSONLogger(os.Stdout)
	logger = log.WithPrefix(logger, "ts", log.DefaultTimestamp)

	err := godotenv.Load(".env")

	if err != nil {
		logger.Log("dotenv", "load", "err", err)
	}


  server := &http.Server{
		Handler:      handler.NewServer(logger),
		Addr:         httpAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Log("http", "server", "addr", httpAddr)

	if err := server.ListenAndServe(); err != nil {
		logger.Log("http", "server", "err", err)
		os.Exit(1)
	}
}

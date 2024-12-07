package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"kreditplus-test/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err.Error())
		}
	}(f)

	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(f)
	fmt.Println("Starting Server ... ", time.Now().Unix())
	startServerTime := time.Now()
	ctx := context.Background()

	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    ":" + viper.GetString("PORT"),
		Handler: server.App(ctx),
	}

	// graceful shutdown
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ... ", time.Since(startServerTime).Seconds(), " s")

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()

	log.Println("Server exiting")
}

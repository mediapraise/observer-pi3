package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"observer-go/src/handler"
	"observer-go/src/router"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	INFO  = "[INFO] "
	ERROR = "[ERROR] "
)

func Create(log *log.Logger) *mux.Router {
	log.Printf("%sServer routes building...", INFO)
	muxRouter := mux.NewRouter()

	(&router.Webhooks{
		Log:     log,
		Router:  muxRouter.PathPrefix("/webhook/").Subrouter(),
		Handler: &handler.Webhook{Log: log},
	}).Build()
	return muxRouter
}

func startServer(log *log.Logger) *http.Server {

	muxRouter := Create(log)
	port := os.Getenv("APP_PORT")
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      muxRouter,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	go func() {
		log.Printf("%sStarting server... in port: %s", INFO, port)
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Printf("%sError starting server: %v", INFO, err)
		}
	}()

	return httpServer
}

// func InitAuth(l *log.Logger) {
// 	refreshFunc := func() (string, time.Duration) {
// 		token, duration := service.NewAuthService(l).Authenticate()
// 		return token, duration
// 	}
// 	auth.RefreshTokenAutomatically(refreshFunc)
// }

func waitForShutdown(s *http.Server, l *log.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <-sigChan
	l.Printf("%sReceived signal: %v", INFO, sig)
	l.Printf("%sReceived terminate, graceful shutdown", INFO)
	d := 30 * time.Second
	tc, cancel := context.WithTimeout(context.Background(), d)
	s.Shutdown(tc)
	defer cancel()
}

func startEnv(log *log.Logger) {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Printf("%sError loading .env file: %v", ERROR, err)
		}
	} else {
		os.Setenv("APP_ENV", "Production")
	}
	log.Printf("%sEnvironment: %s", INFO, os.Getenv("APP_ENV"))
}

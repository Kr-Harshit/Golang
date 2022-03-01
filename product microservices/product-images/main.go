package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-images/files"
	"github.com/Kr-Harshit/golang-example/product-microservices/product-images/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	loadTheEnv()
}

// loadTheEnv loads env variable
func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file; ", err)
	}
}

func main() {
	// logger logs the action on standard output
	logger := log.New(os.Stdout, "products-image", log.LstdFlags)

	// create the storage class, using local storage
	store, err := files.NewLocal(os.Getenv("BASE_PATH"), 1024*1000*5)
	if err != nil {
		logger.Println("Unable to create storage", "error", err)
		os.Exit(1)
	}

	//creating file handler
	fileHandler := handlers.NewFiles(store, logger)

	// creating new ServerMux
	router := mux.NewRouter()

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}}
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fileHandler.Save)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/",
			http.FileServer(http.Dir(os.Getenv("BASE_PATH")))))

	// Serving Docs by middleware on getRouter
	opts := middleware.RedocOpts{SpecURL: "./swagger.yaml"}
	redocHandler := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", redocHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))) //serving swagger.yaml file when client browser reques

	// creating new Server
	serv := http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// context for graceFull termination of server shutdown
	interruptCtx, cancel := context.WithCancel(context.Background())

	// Go Routine to listen to termination signal
	go func() {
		wait := make(chan os.Signal, 1)
		signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

		<-wait
		cancel()
	}()

	// Go Routine to start server
	go func() {
		logger.Printf("Starting server on PORT %s!\n", os.Getenv("PORT"))

		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("[ERROR] starting server:", err)
			os.Exit(1)
		}
	}()

	// Go Routine to shutdown the server on interruption signal
	<-interruptCtx.Done()

	logger.Printf("[INFO] terminating Server")
	if err := serv.Shutdown(interruptCtx); err != nil {
		logger.Println("[ERROR] shutting down Server", err)
	}
	logger.Println("[INFO] Server closed!!!")
}

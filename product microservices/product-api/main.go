package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	protos "github.com/Kr-Harshit/golang-example/product-microservices/currency/protos/currency"
	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/handlers"
	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/models"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[ERROR] loading .env file")
	}
}

func main() {
	log := hclog.Default()
	validator := models.NewValidation()

	// setting gRPC connection
	conn, err := grpc.Dial(os.Getenv("CURRENCY_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Error("Unable to set up TCP connection for gRPC", "error", err)
		panic(err)
	}
	defer conn.Close()

	// create gRPC client for currency
	currencyClient := protos.NewCurrencyClient(conn)

	// create ProductDB
	prodDb := models.NewProductsDB(currencyClient, log)

	// Creating Handlers
	productHandler := handlers.NewProducts(log, validator, prodDb)

	// creating new ServeMux
	router := mux.NewRouter()

	// Creating Method Specific SubRouter and registering handlers
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetProducts).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products", productHandler.GetProducts)

	getRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.GetProduct).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.GetProduct)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.Use(productHandler.ValidateProduct)
	putRouter.HandleFunc("/products", productHandler.Update)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.Use(productHandler.ValidateProduct)
	postRouter.HandleFunc("/products", productHandler.Create)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete)

	// Serving Docs by middleware on getRouter
	opts := middleware.RedocOpts{SpecURL: "./swagger.yaml"}
	redocHandler := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", redocHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))) //serving swagger.yaml file when client browser reques

	// CORS
	allowed_origins_opts := []string{"http://localhost:3000"}
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins(allowed_origins_opts))

	// creating new server
	serv := http.Server{
		Addr:         os.Getenv("SERVER_ADDRESS"),
		Handler:      corsHandler(router),
		ErrorLog:     log.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// context for gracefull termination of server
	ctx, cancel := context.WithCancel(context.Background())

	// creating go routine to listen to termination signal

	go func() {
		wait := make(chan os.Signal, 1)
		signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

		<-wait
		cancel()
	}()

	// go routine to start server
	go func() {
		log.Info("Starting server", "address", serv.Addr)
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Unable to start server", "Address", serv.Addr, "error", err)
			os.Exit(1)
		}
	}()

	// go routine to shutdown the server on recieving singal from os
	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("Shutting Server in 5 seconds")
	if err := serv.Shutdown(ctxShutDown); err != nil {
		log.Error("Unable to gracefully shutdown server", "error", err)
		os.Exit(1)
	}
	log.Info("Server Closed!!!")
}

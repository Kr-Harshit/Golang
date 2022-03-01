/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kr-Harshit/golang-example/product-microservices/currency/data"
	protos "github.com/Kr-Harshit/golang-example/product-microservices/currency/protos/currency"
	"github.com/Kr-Harshit/golang-example/product-microservices/currency/server"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Currency rate conversion server.",
	Long:  `A gRPC server for currency rate conversion.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	loadTheEnv()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// loadTheEnv loads env variable when server cmd from terminal is called.
func loadTheEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[ERROR] loading .env file")
	}
}

const BASE_CURRENCY = "USD"

// run starts and serve gRPC server
func run() {
	log := hclog.Default()
	rates, err := data.NewRates(log, BASE_CURRENCY)
	if err != nil {
		log.Error("Unable to fetch rates", "error", err)
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Error("Unable to Listen on port", os.Getenv("SERVER_PORT"), "error", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()               //gRPC Server
	cs := server.NewCurrency(log, rates) //  Currency Server Service

	// Registering Currency Services to gRPC server
	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	ctx, cancel := context.WithCancel(context.Background())

	// Go Routine to Listen for termination Command
	go func() {
		wait := make(chan os.Signal, 1)
		// listening for terminaation signal and passing to wait channel
		signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

		<-wait
		cancel()
	}()

	// Go Routine to Start gRPC server.
	go func() {
		log.Info(fmt.Sprintf("Starting Server at %s", lis.Addr()))

		// Serving gRPC server with TCP connection at mentioned port
		if err := gs.Serve(lis); err != nil {
			log.Error("Unable to start gRPC server", "error", err)
		}
	}()

	<-ctx.Done()
	log.Info("Recieved Termination Signal, Shutting down Server")
	gs.GracefulStop()
	log.Info("Server Closed")

}

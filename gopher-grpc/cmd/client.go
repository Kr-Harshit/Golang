/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"os"
	"time"

	gopher_GRPC "github.com/Kr_Harshit/golang/gopher_GRPC/pkg/gopher"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Query the gRPC server.",
	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("[ERROR] connecting: %s\n", err)
		}
		defer conn.Close()

		client := gopher_GRPC.NewGopherClient(conn)

		var name string

		// Contact the server and print out its response.
		// name := defaultName
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := client.GetGopher(ctx, &gopher_GRPC.GopherRequest{Name: name})
		if err != nil {
			log.Fatalf("[ERROR] could not greet: %v", err)
		}
		log.Printf("[INFO] URL: %s", response.GetMessage())
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const (
	address     = "localhost:9000"
	defaultName = "dr-who"
)

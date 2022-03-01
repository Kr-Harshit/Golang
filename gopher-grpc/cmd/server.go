/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	gopher_grpc "github.com/Kr_Harshit/golang/gopher_GRPC/pkg/gopher"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts The schema gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		listen, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("[ERROR] Failed to Listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		// Register Service
		gopher_grpc.RegisterGopherServer(grpcServer, &Server{})

		log.Printf("[INFO] GRPC server listening on %v", listen.Addr())

		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const (
	port         = ":9000"
	KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"
)

// Sever is used to implement gopher.GopherServer
type Server struct {
	gopher_grpc.UnimplementedGopherServer
}

type Gopher struct {
	URL string `json:"url"`
}

func (s *Server) GetGopher(ctx context.Context, req *gopher_grpc.GopherRequest) (*gopher_grpc.GopherReply, error) {
	res := &gopher_grpc.GopherReply{}

	// Check Request
	if req == nil {
		fmt.Println("[ERROR] request must not be nil")
		return res, xerrors.Errorf("request must not be nil")
	}

	if req.Name == "" {
		fmt.Println("[ERROR] Name mmust not be empty in request")
		return res, xerrors.Errorf("name must not be empty in request")
	}

	log.Printf("Received: %v", req.GetName())

	// Call KutoGo API in order to get Gopher URL
	response, err := http.Get(KuteGoAPIURL + "/gopher?name=" + req.GetName())
	if err != nil {
		log.Fatalf("[ERROR] failed to call KuteGoAPI: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		// Transform our response to []Byte
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("[ERROR] Failed to read response body: %v", err)
		}

		// Put only needed informations of Json document in our array of Gopher
		var data []Gopher
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Fatalf("[ERROR] Failed to Unmarshal JSON: %v", err)
		}

		// Create a string with all of the Gopher's name
		// Add blank Line as seperator
		var gophers strings.Builder
		for _, gopher := range data {
			gophers.WriteString(gopher.URL + "\n")
		}

		res.Message = gophers.String()
	} else {
		log.Fatal("[ERROR] Can't get the gopher :-(")
	}

	return res, nil
}

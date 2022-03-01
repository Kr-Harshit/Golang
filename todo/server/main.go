package main

import (
	"fmt"
	"log"
	"os"

	"githhub.com/Kr-Harshit/golang-react-todo/router"
)

// func init() {
// 	godotenv.Load()
// }

func getPORT() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to ", port)
	}
	return ":" + port
}

func main() {
	PORT := getPORT()
	r := router.Router()
	fmt.Println("Starting server on port ", PORT)

	if err := r.Run(PORT); err != nil {
		log.Fatal("ERROR: Starting server, ", err)
	}
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This command will call Github repository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		gopherName := "dr-who"

		if len(args) > 0 && args[0] != "" {
			gopherName = args[0]
		}
		getGopher("https://github.com/scraly/gophers/raw/main/", gopherName)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getGopher(baseURI, gopherName string) {
	URL := baseURI + gopherName + ".png"

	fmt.Printf("[INFO] fetching '%s' Gopher...\n", URL)

	// Get the data
	response, err := http.Get(URL)
	if err != nil {
		fmt.Printf("[ERROR] unable to fetch data - %s\n", err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		// Create the file
		filename := gopherName + ".png"
		out, err := os.Create(filename)
		if err != nil {
			fmt.Printf("[ERROR] unable to create file '%s'- %s\n", filename, err)
		}
		defer out.Close()

		// Write the data to file
		_, err = io.Copy(out, response.Body)
		if err != nil {
			fmt.Printf("[ERROR] unable to write data to file '%s'- %s\n", filename, err)
		}

		absolutePath, err := filepath.Abs(filename)
		if err != nil {
			fmt.Printf("[ERROR] File '%s' created but unable to fetch file path:- '%s'\n", filename, err)
		}
		fmt.Printf("[SUCCESS] Data saved at: %s", absolutePath)
	} else {
		fmt.Printf("[ERROR] File '%s'.png doesn't exist'\n", gopherName)
	}
}

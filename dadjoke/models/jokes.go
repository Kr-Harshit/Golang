package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func GetRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responsebytes := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(responsebytes, &joke); err != nil {
		log.Printf("[ERROR] Could not deserialize joke - %v\n", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(http.MethodGet, baseAPI, nil)
	if err != nil {
		log.Printf("[ERROR] Could not request a dadjoke - %v\n", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/Kr-Harshit/golang/dadjokes)")

	reponse, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("[ERROR] Request failed - %v\n", err)
	}

	responseBytes, err := ioutil.ReadAll(reponse.Body)
	if err != nil {
		log.Printf("[ERROR] Could not read response body - %v\n", err)
	}

	return responseBytes
}

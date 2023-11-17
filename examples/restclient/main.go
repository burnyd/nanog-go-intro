package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// General responsestruct for marshalling the response of the API
type ResponseSrtuct struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	// Make the http request
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//Read the response body
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Cannot marshall data", err)
	}
	// Create a pointer to the response struct
	JsonResponse := &ResponseSrtuct{}
	//Unmarshall it 
	err = json.Unmarshal(responseData, &JsonResponse)
	// Show how type values work.
	fmt.Printf("This joke is hilairous: %s\n", JsonResponse.Joke)
}

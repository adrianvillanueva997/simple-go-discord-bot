package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type catResponse struct {
	PhotoUrl string   `json:"url"`

}
type catResponses []catResponse

func getCatPhoto() string {
	log.Println("Connecting to Cat Photos")
	response, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Println(string(responseData))
	var catResponse catResponses
	err = json.Unmarshal(responseData, &catResponse)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return catResponse[0].PhotoUrl

}

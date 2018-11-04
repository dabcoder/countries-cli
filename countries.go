package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Country struct {
	Capital string
	//Currency string `json:"`
	//Language string
}

type countryResp struct {
	Collection []Country
}

func main() {
	const baseURL = "https://restcountries.eu/rest/v2"
	countryName := os.Args[1]
	flag.Parse()
	resp, err := http.Get(baseURL + "/name/" + countryName + "?fullText=true")
	if err != nil {
		fmt.Println("There was an issue completing this request")
		log.Fatal(err)
	}
	fmt.Println(baseURL + "/name/" + countryName + "?fullText=true")
	fmt.Println(resp)

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	countries := make([]Country, 0)

	json.Unmarshal(body, &countries)

	fmt.Println("The capital of "+countryName+" is: ", countries)
}

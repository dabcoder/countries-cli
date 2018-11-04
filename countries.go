package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type country struct {
	Capital string `json:"capital"`
	//Currency string `json:"`
	//Language string
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

	bodyText := string(body)

	fmt.Println(bodyText)
}

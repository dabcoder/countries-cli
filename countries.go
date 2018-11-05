package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Country struct {
	Capital  string         `json:"capital"`
	Currency []CurrencyJson `json:"currencies"`
	Language []LanguageJson `json:"languages"`
}

type CurrencyJson struct {
	CurrencyName   string `json:"name"`
	CurrencySymbol string `json:"symbol"`
}

type LanguageJson struct {
	LanguageName string `json:"name"`
}

func main() {
	const baseURL = "https://restcountries.eu/rest/v2"
	var fullURL string
	countryName := os.Args[1:]
	if len(countryName) > 1 {
		countryWeb := strings.Join(countryName, "%20")
		fullURL = baseURL + "/name/" + countryWeb + "?fullText=true"
	} else {
		fullURL = baseURL + "/name/" + countryName[0] + "?fullText=true"
	}
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("There was an issue completing this request")
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	countries := make([]Country, 0)
	json.Unmarshal(body, &countries)

	fmt.Println("Capital:", countries[0].Capital)
	fmt.Println("Currency:", countries[0].Currency[0].CurrencyName)
	fmt.Println("Symbol:", countries[0].Currency[0].CurrencySymbol)
	fmt.Println("Language:", countries[0].Language[0].LanguageName)
}

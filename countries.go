package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli"
)

const baseURL = "https://restcountries.eu/rest/v2/name/"
const endOfURL = "?fullText=true"

type Country struct {
	Capital    string         `json:"capital"`
	Subregion  string         `json:"subregion"`
	Population int            `json:"population"`
	Currency   []CurrencyJson `json:"currencies"`
	Language   []LanguageJson `json:"languages"`
}

type CurrencyJson struct {
	CurrencyName   string `json:"name"`
	CurrencySymbol string `json:"symbol"`
}

type LanguageJson struct {
	LanguageName string `json:"name"`
}

func main() {
	var fullURL string
	var countryName string

	app := cli.NewApp()
	app.Name = "countries of the world CLI"
	app.Version = "0.1.0"
	app.Usage = "info about a given country"

	// To display the country's population or language only in the output
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "population, p",
			Usage: "Display the country's population only",
		},
		cli.BoolFlag{
			Name:  "language, l",
			Usage: "Display the country's language only",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			return cli.NewExitError("Needs a country", 0)
		} else if c.NArg() > 1 {
			countryName = strings.Join(c.Args(), "%20")
			fullURL = fmt.Sprintf("%s%s%s", baseURL, countryName, endOfURL)
		} else {
			countryName = c.Args()[0]
			fullURL = fmt.Sprintf("%s%s%s", baseURL, countryName, endOfURL)
		}

		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Println("There was an issue completing this request")
			log.Fatal(err)
		}
		// Mispelled country name or response code != 200
		if resp.StatusCode != 200 {
			return cli.NewExitError("Couldn't find any info", 0)
		}
		defer resp.Body.Close()

		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		countries := make([]Country, 0)
		json.Unmarshal(body, &countries)

		// Popluation or Language flag has been enabled
		if c.Bool("population") {
			fmt.Println("Population:", countries[0].Population)
		} else if c.Bool("language") {
			fmt.Println("Language:", countries[0].Language[0].LanguageName)
		} else {
			fmt.Println("Capital:", countries[0].Capital)
			fmt.Println("Region:", countries[0].Subregion)
			fmt.Println("Population:", countries[0].Population)
			fmt.Println("Currency:", countries[0].Currency[0].CurrencyName, countries[0].Currency[0].CurrencySymbol)
			fmt.Println("Language:", countries[0].Language[0].LanguageName)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

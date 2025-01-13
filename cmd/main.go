package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	pokemon "github.com/codescalersinternships/PokeAPIHTTPClient-RawanMostafa/pkg"
)

const defaultEndpoint = "/machine"

func getFlags() string {

	var endpoint string
	flag.StringVar(&endpoint, "endpoint", "", "Specifies the endpoint")

	flag.Parse()
	return endpoint
}

func decideConfigs() string {

	endpoint := getFlags()

	if endpoint == "" {
		envEndpoint, found := os.LookupEnv("ENDPOINT")

		if found {
			endpoint = envEndpoint
		} else {
			endpoint = defaultEndpoint
		}
	}
	return endpoint

}

func main() {
	endpoint := decideConfigs()

	c := pokemon.NewClient(endpoint, time.Second)
	pokemon, err := c.Pokemon(1)
	if err != nil {
		log.Fatal(err)
	}
	poke, _ := json.MarshalIndent(pokemon, "", "		")
	log.Println(string(poke))
}

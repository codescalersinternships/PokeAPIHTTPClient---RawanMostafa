package main

import (
	"flag"
	"log"
	"os"
	"time"

	pkg "github.com/codescalersinternships/PokeAPIHTTPClient-RawanMostafa/pkg"
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

	c := pkg.NewClient(endpoint, time.Second)
	resource, err := c.Resource()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", resource)
}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	pkg "github.com/codescalersinternships/PokeAPIHTTPClient-RawanMostafa/pkg"
)

const defaultEndpoint = "/machine"

func readBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error in reading request body: %v", err)

	}
	return string(body), nil
}

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
	resp, err := c.RetrySendRequest()
	if err != nil {
		log.Fatal(err)
	}
	body, err := readBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
}

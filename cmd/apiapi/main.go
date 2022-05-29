package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/datoga/apiapi"
)

func main() {
	link, err := url.Parse("https://www.example.com")

	if err != nil {
		panic(err)
	}

	result, err := apiapi.ApiApiFromURL(*link)

	if err != nil {
		panic(err)
	}

	bytes, err := json.MarshalIndent(result, "", "")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

package apiapi

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func ApiApiFromURL(link url.URL, options ...Option) (Result, error) {
	response, err := http.Get(link.String())

	if err != nil {
		return Result{}, fmt.Errorf("failed getting data from link %v (%v)", link, err)
	}

	if response.StatusCode != http.StatusOK {
		return Result{}, fmt.Errorf("response not 200, got %v", response.StatusCode)
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println("failure closing body with error", err)
		}
	}()

	result, err := ApiApi(response.Body, options...)

	if err != nil {
		return Result{}, err
	}

	result.Network.Headers = response.Header

	return result, nil
}

func ApiApi(reader io.Reader, options ...Option) (Result, error) {
	bHTML, err := ioutil.ReadAll(reader)

	if err != nil {
		return Result{}, fmt.Errorf("failed reading from reader with error %v", err)
	}

	encoded := base64.StdEncoding.EncodeToString(bHTML)

	result := Result{
		HTML: encoded,
	}

	return result, nil
}

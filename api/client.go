package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emc-dojo/govmaxapi/model"
)

type Client struct {
	host       string
	username   string
	password   string
	insecure   bool
	httpClient http.Client
}

func NewClient(host, username, password string, insecure bool) (Client, error) {
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		},
	}

	return Client{host, username, password, insecure, httpClient}, nil
}

func (apiClient Client) APICall(httpMethod, uri string, createParam interface{}) error {
	b, err := json.Marshal(createParam)
	if err != nil {
		return fmt.Errorf("error marshalling param %s\n", err)
	}

	req, err := http.NewRequest(httpMethod, uri, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("error creating http request %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(apiClient.username, apiClient.password)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making http call %s\n", err)
	}
	defer resp.Body.Close()

	var genericResult model.GenericResult
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(genericResult)
	if err != nil {
		return fmt.Errorf("error reading create json response %s\n", err)
	}

	if !genericResult.Success {
		return fmt.Errorf("error creating %s\n", genericResult.Message)
	}

	return nil
}

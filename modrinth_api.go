package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type ModrinthClient struct {
	httpClient http.Client
}

func NewModrinthClient() ModrinthClient {
	out := ModrinthClient{}
	return out
}

func (mc *ModrinthClient) makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "hibiii/bluemon/no-version")

	resp, err := mc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("makeRequest: bad status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

const version_route = "https://api.modrinth.com/v2/version/"

func (mc *ModrinthClient) GetVersionDownloads(version string) (int, error) {
	body, err := mc.makeRequest(version_route + version)
	if err != nil {
		return 0, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}
	downloads_field := data["downloads"]
	if downloads_field == nil || reflect.TypeOf(downloads_field).Kind() != reflect.Float64 {
		return 0, fmt.Errorf("ModrinthClient.GetVersionDownloads: could not find downloads field! found %q instead", downloads_field)
	}
	downloads := int(downloads_field.(float64))

	return downloads, nil
}

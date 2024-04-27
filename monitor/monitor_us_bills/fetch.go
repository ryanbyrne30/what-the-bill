package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Fetch struct {
	apiKey string
	mongo  *Mongo
}

func NewFetch(apiKey string, mongo *Mongo) *Fetch {
	return &Fetch{apiKey: apiKey, mongo: mongo}
}

type Package struct {
	PackageID   string `json:"packageId"`
	UpdatedAt   string `json:"lastModified"`
	PackageLink string `json:"packageLink"`
}

type Response struct {
	Packages []Package `json:"packages"`
}

func (f *Fetch) BillsSince(since time.Time, limit, offset int) ([]Package, error) {
	d := since.Format("2006-01-02T15:04:05Z")
	url := fmt.Sprintf("https://api.govinfo.gov/collections/BILLS/%s?pageSize=%d&offset=%d", d, limit, offset)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Api-Key", f.apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data Response

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	log.Printf("Fetched %d packages", len(data.Packages))
	return data.Packages, nil
}

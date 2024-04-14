package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Fetch struct {
	apiKey string
}

func NewFetch(apiKey string) *Fetch {
	return &Fetch{apiKey: apiKey}
}

type FetchBillSummaryShortTitle struct {
	Title string `json:"title"`
}

type FetchBillsSummaryMember struct {
	Role    string `json:"role"`
	Chamber string `json:"chamber"`
	BioID   string `json:"bioGuideId"`
	Name    string `json:"memberName"`
	State   string `json:"state"`
	Party   string `json:"party"`
}

type FetchBillSummaryCommittee struct {
	AuthorityID string `json:"authorityId"`
	Chamber     string `json:"chamber"`
	Name        string `json:"committeeName"`
	Type        string `json:"type"`
}

type FetchBillSummaryDownload struct {
	TextLink string `json:"txtLink"`
}

type FetchBillSummaryResponse struct {
	ShortTitle  []FetchBillSummaryShortTitle `json:"shortTitle"`
	Title       string                       `json:"title"`
	Url         string                       `json:"detailsLink"`
	Congress    string                       `json:"congress"`
	Session     string                       `json:"session"`
	Pages       string                       `json:"pages"`
	Version     string                       `json:"billVersion"`
	Type        string                       `json:"billType"`
	Members     []FetchBillsSummaryMember    `json:"members"`
	Committees  []FetchBillSummaryCommittee  `json:"committees"`
	Download    FetchBillSummaryDownload     `json:"download"`
	PublishedAt string                       `json:"dateIssued"`
	UpdatedAt   string                       `json:"lastModified"`
}

func (f *Fetch) Sleep() {
	time.Sleep(2 * time.Second)
}

func (f *Fetch) FetchBillDetails(url string) (*FetchBillSummaryResponse, error) {
	f.Sleep()
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

	var data FetchBillSummaryResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	log.Printf("Fetched summary for %s", url)
	return &data, nil
}

func (f *Fetch) FetchBillText(url string) (string, error) {
	f.Sleep()
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

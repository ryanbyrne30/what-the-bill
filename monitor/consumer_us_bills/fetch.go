package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

type FetchBillSummaryRelated struct {
	BillStatusLink string `json:"billStatusLink"`
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
	Related     FetchBillSummaryRelated      `json:"related"`
	Members     []FetchBillsSummaryMember    `json:"members"`
	Committees  []FetchBillSummaryCommittee  `json:"committees"`
	Download    FetchBillSummaryDownload     `json:"download"`
	PublishedAt string                       `json:"dateIssued"`
	UpdatedAt   string                       `json:"lastModified"`
}

type FetchBillStatusBillAction struct {
	XMLName    xml.Name `xml:"item"`
	ActionDate string   `xml:"actionDate"`
	Text       string   `xml:"text"`
}

type FetchBillStatusBillActions struct {
	XMLName xml.Name                    `xml:"actions"`
	Items   []FetchBillStatusBillAction `xml:"item"`
}

type FetchBillStatusBill struct {
	XMLName xml.Name                   `xml:"bill"`
	Actions FetchBillStatusBillActions `xml:"actions"`
}

type FetchBillStatusResponse struct {
	XMLName xml.Name            `xml:"billStatus"`
	Bill    FetchBillStatusBill `xml:"bill"`
}

func (f *Fetch) Sleep() {
	time.Sleep(2 * time.Second)
}

func (f *Fetch) FetchBillDetails(url string) (*FetchBillSummaryResponse, error) {
	f.Sleep()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if f.apiKey == "" {
		log.Printf("WARN no api key set")
	}
	req.Header.Set("X-Api-Key", f.apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Printf("ERROR received status code: %d. %s", res.StatusCode, res.Body)
		return nil, fmt.Errorf("received bad status code: %d", res.StatusCode)
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
	log.Printf("Fetching bill text from: %s", url)
	f.Sleep()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if f.apiKey == "" {
		log.Printf("WARN no api key set")
	}
	req.Header.Set("X-Api-Key", f.apiKey)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		log.Printf("ERROR received status code when fetching text: %d. %s", res.StatusCode, res.Body)
		return "", fmt.Errorf("received bad status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read bill text from: %s. %v", url, err)
		return "", err
	}

	log.Printf("Fetched text from: %s", url)
	return string(body), nil
}

func (f *Fetch) FetchBillActions(url string) (*FetchBillStatusResponse, error) {
	log.Printf("Fetching bill actions from: %s", url)
	f.Sleep()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if f.apiKey == "" {
		log.Printf("WARN no api key set")
	}
	req.Header.Set("X-Api-Key", f.apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Printf("ERROR received status code when fetching bill actions: %d. %s", res.StatusCode, res.Body)
		return nil, fmt.Errorf("received bad status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read response from fetching bill actions")
		return nil, err
	}

	var response FetchBillStatusResponse
	err = xml.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Could not unmarshal response from fetching bill actions")
		return nil, err
	}

	log.Printf("Fetched bill actions from: %s", url)
	return &response, nil
}

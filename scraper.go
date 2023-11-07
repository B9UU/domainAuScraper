package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func Reqeust(db *sql.DB, page int) {
	defer Timer(page)()
	url := fmt.Sprintf("https://www.domain.com.au/sale/dianella-wa-6059?excludeunderoffer=1&page=%v", page)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authority", "www.domain.com.au")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "en-US,en;q=0.5")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Something went wrong", err)
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	var responseData ResponseType

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&responseData)
	if err != nil {
		fmt.Println("Failed to decode JSON response:", err)
	}
	for _, v := range responseData.Props.ListingsMap {
		// fmt.Println(k, v.ListingModel.Address)
		_, err = insertRow(db, "first", v)
		if err != nil {
			log.Fatalln("Something went wrong insering the row", err)
		}
		// fmt.Println(k, "to db success")
	}

}

type Requester struct {
	client    *http.Client
	rateLimit <-chan time.Time
	wg        sync.WaitGroup
}

// NewRequester creates a new Requester with the specified request per second rate.
func NewRequester(requestsPerSecond int) *Requester {
	return &Requester{
		client:    &http.Client{},
		rateLimit: time.Tick(time.Second / time.Duration(requestsPerSecond)),
	}
}

// SendRequest sends an HTTP request to the given URL and waits for it to complete.
func (r *Requester) SendRequest(db *sql.DB, page int) {
	<-r.rateLimit // Wait for the rate limiter to allow the next request
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		Reqeust(db, page)
		// Process the response as needed
		// For example, you can read the response body, check the status code, etc.
		// fmt.Println("Response Status:", resp.Status)
	}()
}

// Wait waits for all the requests to complete.
func (r *Requester) Wait() {
	r.wg.Wait()
}

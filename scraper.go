package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func parsing(resp *http.Response) {

}

func Reqeust() ResponseType {

	url := "https://www.domain.com.au/sale/dianella-wa-6059?excludeunderoffer=1&lastsearchdate=2023-11-04t20%3A45%3A58.081z&page=2"

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
	return responseData

}

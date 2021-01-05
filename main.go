package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Price 最終価格
type Price struct {
	LastPrice float64 `json:"last_price"`
}

func main() {
	uri := "https://api.zaif.jp/api/1/last_price/xem_jpy"
	price, err := getLastPrice(uri)
	if err != nil {
		log.Fatal(err)
	}
	message := fmt.Sprintf("XEM最終価格（Go）%.2f JPY", price.LastPrice)
	fmt.Println(message)
	sendlinemessage(message)
}

func getLastPrice(url string) (*Price, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// fmt.Println(string(bytes))

	var price Price
	err = json.Unmarshal(bytes, &price)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// fmt.Printf("%.2f\n", price.LastPrice)

	return &price, nil
}

func sendlinemessage(message string) {
	token := "Your access token"
	api := "https://notify-api.line.me/api/notify"

	values := url.Values{}
	values.Add("message", message)

	req, err := http.NewRequest("POST", api, strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("メッセージ送信結果: %s\n", body)
}

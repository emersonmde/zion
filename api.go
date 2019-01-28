package zion

import (
	"log"
	"net/url"
)

type Client struct {
	homeServerUrl *url.URL
	accessToken   string
}

func NewClient(homeServerUrl, accessToken string) *Client {
	url, err := url.Parse(homeServerUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		homeServerUrl: url,
		accessToken:   accessToken,
	}
}



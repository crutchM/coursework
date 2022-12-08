package main

import (
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{client: &http.Client{
		Timeout: 5 * time.Second}}
}

func (s *Client) SendRequestToMainService(repository GithubRepository) {

}

func (s *Client) SendRequestToGithubApi(url string) {

}

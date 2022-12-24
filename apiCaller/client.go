package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
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
	resp, err := s.client.Get(fmt.Sprint(viper.GetString("base_url"), "/", url))
	if err != nil {
		log.Print(err)
		return
	}
	resp.Body.Close()
	var repo GithubRepository
	json.NewDecoder(resp.Body).Decode(&repo)
	s.SendRequestToMainService(repo)
}

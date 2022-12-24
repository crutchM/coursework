package main

import (
	bytes "bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	client *http.Client
}

func GenerateClient() *Client {
	return &Client{client: &http.Client{
		Timeout: 5 * time.Second}}
}

func (s *Client) SendRequestToMainService(repository GithubRepository) {
	jsonStr, err := json.Marshal(repository)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", fmt.Sprint(viper.GetString("main_service_url"), "/data"), bytes.NewBuffer(jsonStr))
	_, err = s.client.Do(req)
	if err != nil {
		log.Println(err)
	}
	//defer resp.Body.Close()
}

func (s *Client) SendRequestToGithubApi(url string) {
	var tmp = strings.ReplaceAll(url, "https://", "")
	var arr = strings.Split(tmp, "/")
	var requ = fmt.Sprint(viper.GetString("base_url"), "/", arr[1], "/", arr[2])
	resp, err := s.client.Get(requ)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	var repo GithubRepository
	fmt.Println(json.NewDecoder(resp.Body))
	err = json.NewDecoder(resp.Body).Decode(&repo)
	if err != nil {
		log.Println(err)
	}
	s.SendRequestToMainService(repo)
}

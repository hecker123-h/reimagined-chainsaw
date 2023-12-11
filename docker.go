package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	registryURL = "301522045249.dkr.ecr.ap-south-1.amazonaws.com"
	username    = "hecker" // Replace with your registry username
	password    = "hacker" // Replace with your registry password
)

// RepositoriesResponse represents the response for the list of repositories
type RepositoriesResponse struct {
	Repositories []string `json:"repositories"`
}

func main() {
	client := &http.Client{}

	// Get the list of repositories
	repos, err := getRepositories(client)
	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		fmt.Println("Repository:", repo)
		// Add logic to fetch images/tags for each repository
	}
}

func getRepositories(client *http.Client) ([]string, error) {
	req, err := http.NewRequest("GET", registryURL+"_catalog", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var reposResponse RepositoriesResponse
	if err := json.Unmarshal(body, &reposResponse); err != nil {
		return nil, err
	}

	return reposResponse.Repositories, nil
}

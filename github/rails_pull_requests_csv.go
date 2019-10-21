package main

import (
	"encoding/json"
	"fmt"
	"go_scripts/github/structs"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	railsPullURL := "https://api.github.com/repos/rails/rails/pulls"
	accessToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")

	req, _ := http.NewRequest("GET", railsPullURL, nil)

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()

	res, _ := client.Do(req)

	pullsData, _ := ioutil.ReadAll(res.Body)

	var pullRequests []structs.PullRequest

	json.Unmarshal(pullsData, &pullRequests)

	for i := 0; i < len(pullRequests); i++ {
		fmt.Println(pullRequests[i].Title)

		labels := pullRequests[i].Labels

		for j := 0; j < len(labels); j++ {
			fmt.Println(labels[j].Name)
		}
	}
}

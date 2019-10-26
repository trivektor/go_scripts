package main

import (
  "os"
  "net/http"
  "fmt"
  "strconv"
  "go_scripts/github/structs"
  "io/ioutil"
  "encoding/json"
)

// https://stackoverflow.com/questions/2439453/using-a-pointer-to-array
func PrintlnIssues(issues []structs.Issue) {
  for i := 0; i < len(issues); i++ {
    issue := issues[i]
    fmt.Printf("[%d] %s\n", issue.Number, issue.Title)
  }
}

func main ()  {
  client := &http.Client{}
  accessToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
  railsIssuesURL := "https://api.github.com/repos/rails/rails/issues"

  req, _ := http.NewRequest("GET", railsIssuesURL, nil)
  page := 1

  for true {
    q := req.URL.Query()
    q.Add("access_token", accessToken)
    q.Add("page", strconv.Itoa(page))
    req.URL.RawQuery = q.Encode()

    res, _ := client.Do(req)

    issuesData, _ := ioutil.ReadAll(res.Body)

    var issues[]structs.Issue

    json.Unmarshal(issuesData, &issues)

    if len(issues) == 0 {
      break
    }

    PrintlnIssues(issues)

    page += 1
  }
}

package main

import (
  "go_scripts/github/constants"
  "net/http"
  "fmt"
  "io/ioutil"
  "log"
  "encoding/json"
)

type Actor struct {
  Id uint64 `json:"actor"`
  Login string `json:"login"`
  DisplayLogin string `json:"display_login"`
}

type Event struct {
  Id uint64 `json:"id"`
  Type string `json:"type"`
  Actor Actor `json:"actor"`
}

type Events struct {
  Collection []Event
}

func main() {
  url := constants.GITHUB_API_URL + "/events"
  fmt.Println(url)
  res, err := http.Get(url)

  if (err != nil) {
    log.Fatalln(err)
  }

  byteValue, err := ioutil.ReadAll(res.Body)

  if (err != nil) {
    log.Fatalln(err)
  }

  var events []map[string]interface{}
  //var events []Event
  json.Unmarshal(byteValue, &events)
  mappings := make(map[string]interface{})

  var eventsOfType []interface{}

  for i := 0; i < len(events); i++ {
    //var actor Object
    //json.Unmarshal(events[i]["actor"], &actor)
//    fmt.Println(events[i])
    //actor := events[i]["actor"].(map[string]interface{})
    eventType := events[i]["type"].(string)

    if _, ok := mappings[eventType]; ok {
      //eventsList := make([]interface{}, 0)
      eventsOfType = append(eventsOfType, events[i])
    } else {
      mappings[eventType] = eventsOfType
    }
  }

  keys := make([]string, len(mappings))

  i := 0

  for k := range mappings {
    keys[i] = k
    i++
  }

  fmt.Println(keys)
  fmt.Println(keys[0])
  firstKey := keys[0]

  _eventsOfType := mappings[firstKey].([]interface{})

  fmt.Println(_eventsOfType)

  defer res.Body.Close()
}

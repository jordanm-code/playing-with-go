package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

func main() {
  fmt.Println("Starting up HTTP...")

  xchgURL := "https://api.exchangerate.host/latest"

  client := &http.Client{
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    },
  }

  resp, err := client.Get(xchgURL)
  if err != nil {
    log.Fatal(err)
  }

  if resp.StatusCode == http.StatusOK {
    var result map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&result)
    fmt.Println(result["rates"])
  }
}

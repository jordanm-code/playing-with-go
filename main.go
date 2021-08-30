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

  client := &http.Client{}

  resp, err := client.Get(xchgURL)
  if err != nil {
    log.Fatal(err)
  }

  if resp.StatusCode == http.StatusOK {
    var result map[string]interface{}

    fmt.Println("Response:")
    fmt.Println(resp)

    fmt.Println("Body")
    fmt.Println(resp.Header)
    fmt.Println(resp.Body)

    json.NewDecoder(resp.Body).Decode(&result)
    fmt.Println(result["rates"])
  }
}

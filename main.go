package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
  Water int `json:"water"`
  Wind int  `json:"wind"`
}

func main() {
  
  const (
    URL = "https://jsonplaceholder.typicode.com/posts"
    iter = 10
  )

  var (
    rnd, rnd2 int
    req *http.Request
    res *http.Response
    client *http.Client
    json_req []byte

    status_water, status_wind string

    result Response
    err error
  )

  data := make(map[string]int)

  for i := 0; i <= iter; i++ {
    rnd = rand.Intn(100-1)+1
    rnd2 = rand.Intn(100-1)+1

    data["water"] = rnd
    data["wind"] = rnd2

    json_req, err = json.Marshal(data)
    client = &http.Client{Timeout: time.Duration(60) * time.Second}
    if err != nil {
      log.Fatalln(err)
    }

    req, err = http.NewRequest("POST", URL, bytes.NewBuffer(json_req))
    req.Header.Set("Content-type", "application/json")
    if err != nil {
      log.Fatalln(err)
    }

    res, err = client.Do(req)
    if err != nil {
      log.Fatalln(err)
    }

    defer res.Body.Close()
   
    body, err := ioutil.ReadAll(res.Body) 
    if err != nil {
      log.Fatalln(err)
    }
   
    err = json.Unmarshal(body, &result)
    if err != nil {
      log.Fatalln(err)
    }

    fmt.Println(string(body))
    
    if result.Water < 5 {
      status_water = "aman"
    } else if result.Water < 8 && result.Water > 6 {
      status_water = "siaga"
    } else {
      status_water = "bahaya"
    } 

    if result.Wind < 6 {
      status_wind = "aman"
    } else if result.Wind < 15 && result.Wind > 7 {
      status_wind = "siaga"
    } else {
      status_wind = "bahaya"
    }

    fmt.Println("status water:", status_water)
    fmt.Println("status wind:", status_wind,"\n")

  }
}

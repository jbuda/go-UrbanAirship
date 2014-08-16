package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const APP_KEY = "xx"
const MASTER_SECRET = "xx"
const CPN = "AAAA000147028"
const API_URL = "go.urbanairship.com/api/device_tokens"

type Feed struct {
	Device_tokens []Device_info
}

type Device_info struct {
	Device_token string
	Active       string
	Alias        string
}

func main() {
	fmt.Printf("\nDevice Token Lookup\n----------------\n")
	fmt.Printf("App Key : " + APP_KEY + "\nMaster Secret : " + MASTER_SECRET + "\n")
	fmt.Printf("CPN : " + CPN + "\n")
	fmt.Printf("----------------\n")

	load_json()
}

func load_json() {

	urlStr := fmt.Sprintf("https://%v:%v@%v", APP_KEY, MASTER_SECRET, API_URL)
	client := &http.Client{}
	r, _ := http.NewRequest("GET", urlStr, nil)

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)

	var data Feed
	json.Unmarshal(body, &data)

	for i, device := range data.Device_tokens {
		fmt.Printf("%d: token ~ %s alias ~ %s\n", i, device.Device_token, device.Alias)
	}
}

/*
{
   "next_page": "https://go.urbanairship.com/api/device_tokens/?start=07AAFE44CD82C2F4E3FBAB8962A95B95F90A54857FB8532A155DE3510B481C13&limit=2",
   "device_tokens_count": 87,
   "device_tokens": [
      {
         "device_token": "0101F9929660BAD9FFF31A0B5FA32620FA988507DFFA52BD6C1C1F4783EDA2DB",
         "active": false,
         "alias": null,
         "tags": []
      },
      {
         "device_token": "07AAFE44CD82C2F4E3FBAB8962A95B95F90A54857FB8532A155DE3510B481C13",
         "active": true,
         "alias": null,
         "tags": ["tag1", "tag2"]
      }
   ],
   "active_device_tokens_count": 37
}*/

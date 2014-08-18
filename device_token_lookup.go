package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const APP_KEY = "xx"
const MASTER_SECRET = "xx"
const CPN = "xx"
const API_URL = "go.urbanairship.com/api/device_tokens"
const LIMIT = "10000"

var devices []Device_info

type Feed struct {
	Next_page                  string
	Device_tokens_count        int
	Active_device_tokens_count int
	Device_tokens              []Device_info
}

type Device_info struct {
	Device_token string
	Active       string
	Alias        string
	Tags         []string
}

func main() {
	fmt.Printf("\nDevice Token Lookup\n----------------\n")
	fmt.Printf("App Key : " + APP_KEY + "\nMaster Secret : " + MASTER_SECRET + "\n")
	fmt.Printf("CPN : " + CPN + "\n")
	fmt.Printf("----------------\n")

	urlStr := fmt.Sprintf("https://%v:%v@%v?limit=%v", APP_KEY, MASTER_SECRET, API_URL, LIMIT)

	load_json(urlStr)
}

func load_json(urlstring string) {

	client := &http.Client{}
	r, _ := http.NewRequest("GET", urlstring, nil)

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)

	var data Feed
	json.Unmarshal(body, &data)

	fmt.Printf("\nRunning %v", urlstring)
	//fmt.Printf("\nNumber of device tokens : %d\nNumber of active device tokens : %d\nNext page : %s\n\n====\n\n", data.Device_tokens_count, data.Active_device_tokens_count, data.Next_page)

	for i, device := range data.Device_tokens {
		_ = i
		if device.Alias == CPN {
			devices = append(devices, device)
			fmt.Printf("%d\nDevice Token : %s\nAlias : %s\nTags : %s\n\n====\n\n", i, device.Device_token, device.Alias, strings.Join(device.Tags, ","))
		}
	}

	if data.Next_page != "" {
		urlStr := fmt.Sprintf("https://%v:%v@%v?%v", APP_KEY, MASTER_SECRET, API_URL, strings.Split(data.Next_page, "?")[1])
		load_json(urlStr)
	} else {

		fmt.Printf("\nDevice tokens\n")

		for _, el := range devices {
			fmt.Printf("\n%v ~ %v", el.Device_token, el.Alias)
		}

		fmt.Printf("\n\n=========\nCompleted\n=========\n\n")

	}

}

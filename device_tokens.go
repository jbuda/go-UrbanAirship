package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const APP_KEY = "<APP_KEY>"
const MASTER_SECRET = "<MASTER_SECRET>"
const ALIAS = ""
const API_URL = "go.urbanairship.com/api/device_tokens"
const LIMIT = 10000

var devices []Device_info
var counter int = 0
var upper int = 0
var lower int = counter

type Feed struct {
	Next_page                  string
	Device_tokens_count        int
	Active_device_tokens_count int
	Device_tokens              []Device_info
}

type Device_info struct {
	Device_token string
	Active       bool
	Alias        string
	Tags         []string
}

func main() {
	fmt.Printf("\nDevice Tokens\n----------------\n")
	fmt.Printf("App Key : " + APP_KEY + "\nMaster Secret : " + MASTER_SECRET + "\n")
	fmt.Printf("Alias : " + ALIAS + "\n")
	fmt.Printf("----------------\n")

	urlStr := fmt.Sprintf("https://%v:%v@%v?limit=%v", APP_KEY, MASTER_SECRET, API_URL, LIMIT)

	load_json(urlStr)
}

func load_json(urlstring string) {

	counter++
	client := &http.Client{}
	r, _ := http.NewRequest("GET", urlstring, nil)

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)

	var data Feed
	json.Unmarshal(body, &data)

	upper = counter * LIMIT
	fmt.Printf("\nRecords : %d - %d", lower, upper)
	lower = upper + 1

	fmt.Printf("\nRunning %v", urlstring)

	for _, device := range data.Device_tokens {
		if device.Alias == ALIAS || ALIAS == "" {
			devices = append(devices, device)
		}
	}

	if data.Next_page != "" && counter < 5 {
		urlStr := fmt.Sprintf("https://%v:%v@%v?%v", APP_KEY, MASTER_SECRET, API_URL, strings.Split(data.Next_page, "?")[1])
		load_json(urlStr)

	} else {
		file := writeToFile()
		fmt.Printf("\n\n=========\nCompleted\nOutput file : %v\n=========\n\n", file)
	}

}

func writeToFile() (file string) {

	file = fmt.Sprintf("devices_%v.txt", int32(time.Now().Unix()))

	f, _ := os.Create(file)
	f.WriteString("id\tdevice token\talias\ttags\tactive")

	for idx, device := range devices {
		f.WriteString(fmt.Sprintf("\n%d\t%s\t%s\t%s\t%t", idx+1, device.Device_token, device.Alias, strings.Join(device.Tags, ","), device.Active))
	}

	f.Close()

	return file
}

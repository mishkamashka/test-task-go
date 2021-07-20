package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func getBlockByTag(tag string) Block {
	if tag[len(tag)-1] == 10 {
		tag = tag[:len(tag)-1]
	}
request:
	url := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%s&boolean=true&apikey=%s", tag, API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)

	block := Block{}
	if _, ok := data["status"]; ok {
		res := fmt.Sprintf("%s", data["result"])
		if strings.Contains(res, "Max") {
			log.Println(data["result"], "with tag:", tag, ", will try again in 1 sec")
			time.Sleep(1 * time.Second)
			goto request
		} else {
			log.Fatal(data)
			return block
		}
	}
	s := data["result"]
	jsonString, _ := json.Marshal(s)
	json.Unmarshal(jsonString, &block)

	return block
}

func getLastBlockNumber() string {
	url := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=%s", API_KEY)
request:
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var data map[string]string

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &data)

	if _, ok := data["status"]; ok {
		res := fmt.Sprintf("%s", data["result"])
		if strings.Contains(res, "Max") {
			log.Println("Max rate limit reached")
			time.Sleep(1 * time.Second)
			goto request
		} else {
			log.Fatal(data)
			return ""
		}
	}
	return data["result"]
}

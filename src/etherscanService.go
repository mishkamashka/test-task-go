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

func getSampleBlock() Block {
	s := "{\"number\":\"0x10d4f\",\"transactions\":[" +
		"{\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0xbf3403210f9802205f426759947a80a9fda71b1e\",\"value\":\"0x1\"}," +

		"{\"from\":\"0xbf3403210f9802205f426759947a80a9fda23123\"," +
		"\"to\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\",\"value\":\"0x2\"}," +

		"{\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\",\"value\":\"0xa\"}," +

		"{\"from\":\"0xbf3403210f9802205f426759947a80a9fda23123\"," +
		"\"to\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\",\"value\":\"0x100\"}," +

		"{\"from\":\"0xbf3403210f9802205f426759947a80a9fda71b1e\"," +
		"\"to\":\"0xbf3403210f9802205f426759947a80a9fda23123\",\"value\":\"0x10000000000\"}," +

		"{\"from\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\"," +
		"\"to\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\",\"value\":\"0x4\"}," +

		"{\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\",\"value\":\"0x0\"}]}"
	var block Block
	json.Unmarshal([]byte(s), &block)

	return block
}

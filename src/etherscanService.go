package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getBlockByTag(tag string) Block {
	if tag[len(tag)-1] == 10 {
		tag = tag[:len(tag)-1]
	}
request:
	url := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%s&boolean=true&apikey=8FMZ3ADCFS5JFX6JBMTY6P71C2I1UMB43T", tag)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)

	block := Block{}
	if _, ok := data["status"]; ok {
		fmt.Println(data["result"], "\ntag ", tag, "\ntry again in 1 sec")
		//todo probably not the best solution, change
		time.Sleep(1 * time.Second)
		goto request
	}
	s := data["result"]
	jsonString, _ := json.Marshal(s)
	json.Unmarshal(jsonString, &block)

	return block
}

func getSampleBlock() Block {
	s := "{\"difficulty\":\"0x1d95715bd14\",\"hash\":\"0x7eb7c23a5ac2f2d70aa1ba4e5c56d89de5ac993590e5f6e79c394e290d998ba8\"," +
		"\"number\":\"0x10d4f\",\"stateRoot\":\"0xd64a0f63e2c7f541e6e6f8548a10a5c4e49fda7ac1aa80f9dddef648c7b9e25f\"," +
		"\"timestamp\":\"0x55c9ea07\",\"transactions\":[{\"blockNumber\":\"0x10d4f\"," +
		"\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0xbf3403210f9802205f426759947a80a9fda71b1e\",\"value\":\"0x1\"}," +
		"{\"blockNumber\":\"0x10d4f\",\"from\":\"0xd2d5862c001b7ba77970e13c173356bf1a551e2e\"," +
		"\"to\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\",\"value\":\"0x2\"}," +
		"{\"blockNumber\":\"0x10d4f\",\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\",\"value\":\"0x3\"}," +
		"{\"blockNumber\":\"0x10d4f\",\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\",\"value\":\"0x4\"}]}"
	var block Block
	json.Unmarshal([]byte(s), &block)

	return block
}

func getLastBlockNumber() string {
request:
	resp, err := http.Get("https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=8FMZ3ADCFS5JFX6JBMTY6P71C2I1UMB43T")
	if err != nil {
		panic(err)
	}
	var data map[string]string

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &data)

	if _, ok := data["status"]; ok {
		fmt.Println("Max rate limit reached")
		//todo probably not the best solution, change
		time.Sleep(1 * time.Second)
		goto request
	}

	return data["result"]
}

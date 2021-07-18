package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getBlock() Block {
	resp, err := http.Get("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=0xc3981b&boolean=true")
	if err != nil {
		panic(err)
	}
	var block Block

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)

	s := data["result"]
	//err = json.Unmarshal(b, &block)

	jsonString, _ := json.Marshal(s)
	fmt.Print(jsonString)
	json.Unmarshal(jsonString, &block)

	return block
}

func getSampleBlock() Block {
	s := "{\"difficulty\":\"0x1d95715bd14\",\"hash\":\"0x7eb7c23a5ac2f2d70aa1ba4e5c56d89de5ac993590e5f6e79c394e290d998ba8\"," +
		"\"number\":\"0x10d4f\",\"stateRoot\":\"0xd64a0f63e2c7f541e6e6f8548a10a5c4e49fda7ac1aa80f9dddef648c7b9e25f\"," +
		"\"timestamp\":\"0x55c9ea07\",\"transactions\":[{\"blockNumber\":\"0x10d4f\"," +
		"\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0xbf3403210f9802205f426759947a80a9fda71b1e\",\"value\":\"0xaa9f075c200000\"}," +
		"{\"blockNumber\":\"0x10d4f\",\"from\":\"0xd2d5862c001b7ba77970e13c173356bf1a551e2e\"," +
		"\"to\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\",\"value\":\"0xe82f47fbea6c200\"}," +
		"{\"blockNumber\":\"0x10d4f\",\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\",\"value\":\"0xde4289155b36ecc\"}," +
		"{\"blockNumber\":\"0x10d4f\",\"from\":\"0x4458f86353b4740fe9e09071c23a7437640063c9\"," +
		"\"to\":\"0x00000000b7ca7e12dcc72290d1fe47b2ef14c607\",\"value\":\"0x0\"}]}"
	var block Block
	json.Unmarshal([]byte(s), &block)

	return block
}

func getLastBlock() string {
	resp, err := http.Get("https://api.etherscan.io/api?module=proxy&action=eth_blockNumber")
	if err != nil {
		panic(err)
	}
	var data map[string]string

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &data)

	return data["result"]
}
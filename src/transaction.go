package main

type Transaction struct {
	From		string	`json:"from"`
	To			string	`json:"to"`
	Value 		string	`json:"value"`
	BlockNumber	string 	`json:"block_number"`

	//blockHash string
	//gas string
	//gasPrice string
	//hash string
	//input string
	//nonce string
	//r string
	//s string
	//transactionIndex string
	//transactionType string
	//v string
}

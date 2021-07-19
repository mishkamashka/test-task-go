package main

type Block struct {
	Number 		string			`json:"number"`
	Transactions 	[]Transaction		`json:"transactions"`
}

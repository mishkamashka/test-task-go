package main

import (
	"fmt"
	"github.com/holiman/uint256"
	"hash/fnv"
	"strconv"
	"time"
)

// Напиши программу, которая выдаст адрес,
// баланс которого изменился больше остальных
// (по абсолютной величине) за последние 100 блоков.

// что за адрес нужно отслеживать - адрес отправителя, получателя? оба? повторяются ли транзакции для отправителя/получателя?

// изменение по абсолютной величине - это просто сумма всех именений без учета знака?
// т.е. здесь: +4, -3, -2 абсолютное изменение это 9 или -1?

var addresses = make(map[uint32]string)
var balances = make(map[uint32]*uint256.Int)

func main(){
	start := time.Now()

	lastBlockNumber := getLastBlockNumber()
	last, err := strconv.ParseInt(lastBlockNumber, 0, 64)
	if err != nil {
		panic(err)
	}

	for i := last - 99; i <= last; i++ {
		curTag := fmt.Sprintf("0x%x\n", i)

		//need to sleep here or smth 'cause max rate limit hits "5 calls per sec/IP" (c) docs <- for registered accounts, with API key
		//time.Sleep(5 * time.Second)
		curBlock := getBlockByTag(curTag)

		for _, transaction := range curBlock.Transactions {
			bufValue := uint256.NewInt(0)
			curFromAddr := hash(transaction.From)
			curToAddr := hash(transaction.To)
			curValue, err := uint256.FromHex(transaction.Value)
			if err != nil {
				panic(err)
			}

			//fmt.Println("from: ", curFromAddr, " to: ", curToAddr,  " value: ", curValue)

			if curValue.IsZero() {
				continue
			}

			// should addresses be added before or after zero check?
			if _, ok := addresses[curFromAddr]; !ok {
				addresses[curFromAddr] = transaction.From
			}
			if _, ok := addresses[curToAddr]; !ok {
				addresses[curToAddr] = transaction.To
			}

			if _, ok := balances[curFromAddr]; ok {
				bufValue.Add(balances[curFromAddr], curValue)
				balances[curFromAddr] = bufValue
			} else {
				//curValue.Neg(curValue)
				balances[curFromAddr] = curValue
			}

			if _, ok := balances[curToAddr]; ok {
				bufValue.Add(balances[curToAddr], curValue)
				balances[curToAddr] = bufValue
			} else {
				balances[curToAddr] = curValue
			}
		}

	}

	fmt.Println(len(balances))
	//for m,n := range balances {
	//	fmt.Println(m, ":  ", n)
	//}

	elapsed := time.Since(start)
	fmt.Printf("Time %s", elapsed)
}


func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

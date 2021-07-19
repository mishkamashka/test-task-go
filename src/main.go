package main

import (
	"fmt"
	"github.com/holiman/uint256"
	"hash/fnv"
	"log"
	"strconv"
	"time"
)

var addresses = make(map[uint32]string)
var balances = make(map[uint32]*uint256.Int)

// "баланс которого изменился больше остальных (по абсолютной величине) за последние 100 блоков"
// пусть до начала работы есть balance A и B. Дальше происходит:
// A + x1 + x2 + x3 + x4 = A1
// B + y1 + y2 + y3 = B1

// Надо сравнивать A1-A и B1-B или x1+x2+x3+x4 и y1+y2+y3?

func main() {
	start := time.Now()
	bufValue := uint256.NewInt(0)
	log.Println("Program has started\nFetching last blocks...")

	lastBlockNumber := getLastBlockNumber()
	last, err := strconv.ParseInt(lastBlockNumber, 0, 64)
	if err != nil {
		panic(err)
	}

	for i := last - 99; i <= last; i++ {
		curTag := fmt.Sprintf("0x%x\n", i)
		curBlock := getBlockByTag(curTag)

		for _, transaction := range curBlock.Transactions {
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
				// neg and subtraction if abs change is something else
				balances[curFromAddr] = curValue
			}

			if _, ok := balances[curToAddr]; ok {
				bufValue.Add(balances[curToAddr], curValue)
				balances[curToAddr] = bufValue
			} else {
				balances[curToAddr] = curValue
			}
		}

		if (last-100-i)%20 == 0 {
			log.Println("20 blocks processed")
		}
	}

	log.Println("Total addresses count: ", len(balances))

	elapsed := time.Since(start)
	log.Printf("Time %s\n", elapsed)

	bufValue = uint256.NewInt(0)
	var maxAddr uint32 = 0

	bufValue.Cmp(bufValue)

	for address, value := range balances {
		if value.Cmp(bufValue) == 1 {
			bufValue = value
			maxAddr = address
		}
	}

	log.Println("Address with most changes: ", addresses[maxAddr])
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}
	return h.Sum32()
}

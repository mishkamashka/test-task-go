package main

import (
	"fmt"
	"github.com/holiman/uint256"
	"hash/fnv"
	"log"
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
	bufValue := uint256.NewInt(0)

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
	}

	log.Println("Total addresses count: ", len(balances))

	elapsed := time.Since(start)
	log.Printf("Time1 %s\n", elapsed)

	bufValue = uint256.NewInt(0)
	var maxAddr uint32 = 0

	bufValue.Cmp(bufValue)

	for address, value := range balances {
		if value.Cmp(bufValue) == 1 {
			bufValue = value
			maxAddr = address
		}
	}

	log.Println("Most changes: ", addresses[maxAddr])
	elapsed = time.Since(start)
	log.Printf("Time2 %s", elapsed)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

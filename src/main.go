package main

import (
	"fmt"
	"github.com/holiman/uint256"
	"hash/fnv"
	"log"
	"strconv"
)

var addresses = make(map[uint32]string)
var balances = make(map[uint32]*Counter)

func main() {
	log.Println("Program has started\nFetching last blocks...")
	maxAddr := GetMostChangedAddress()
	log.Println("Address with most changed balance: ", addresses[maxAddr])
}

func GetMostChangedAddress() uint32 {
	lastBlockNumber := getLastBlockNumber()
	last, err := strconv.ParseInt(lastBlockNumber, 0, 64)
	if err != nil {
		panic(err)
	}

	var maxAddr uint32 = 0
	for i := last - 19; i <= last; i++ {
		curTag := fmt.Sprintf("0x%x\n", i)
		curBlock := getBlockByTag(curTag)
		//curBlock := getSampleBlock()
		for _, transaction := range curBlock.Transactions {
			curFromAddr := hash(transaction.From)
			curToAddr := hash(transaction.To)
			curValue, err := uint256.FromHex(transaction.Value)
			if err != nil {
				panic(err)
			}

			if curValue.IsZero() {
				continue
			}

			//fmt.Println("from: ", curFromAddr, " to: ", curToAddr, " value: ", curValue)

			if _, ok := addresses[curFromAddr]; !ok {
				addresses[curFromAddr] = transaction.From
			}
			if _, ok := addresses[curToAddr]; !ok {
				addresses[curToAddr] = transaction.To
			}

			if counter, ok := balances[curFromAddr]; ok {
				bufValue1 := *curValue
				if counter.lastSign { //if last operation was addition
					if curValue.Cmp(counter.value) == 1 {
						counter.value.Neg(counter.value)
					} else {
						bufValue1.Neg(curValue)
					}
				}
				bufValue1.Add(counter.value, &bufValue1)
				counter.value = &bufValue1
				counter.lastSign = false
			} else {
				balances[curFromAddr] = NewCounter(*curValue, false)
			}

			if counter, ok := balances[curToAddr]; ok {
				bufValue2 := *curValue
				if !counter.lastSign { //if last operation was subtraction
					if curValue.Cmp(counter.value) == 1 {
						counter.value.Neg(counter.value)
					} else {
						bufValue2.Neg(curValue)
					}
				}
				bufValue2.Add(counter.value, &bufValue2)
				counter.value = &bufValue2
				counter.lastSign = true
			} else {
				balances[curToAddr] = NewCounter(*curValue, true)
			}
		}

		//	if (last-100-i)%20 == 0 {
		//		log.Println("20 blocks processed")
		//	}
		//}
		log.Println("Total addresses count: ", len(balances))

		//for m, n := range balances {
		//	fmt.Println(m, ": ", n.value)
		//}

		bufValue := uint256.NewInt(0)
		for address, counter := range balances {
			if counter.value.Cmp(bufValue) == 1 {
				bufValue = counter.value
				maxAddr = address
			}
		}
	}
	return maxAddr
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}
	return h.Sum32()
}

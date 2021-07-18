package main

import (
	"fmt"
	"github.com/holiman/uint256"
	"hash/fnv"
)

// Напиши программу, которая выдаст адрес,
// баланс которого изменился больше остальных
// (по абсолютной величине) за последние 100 блоков.


// map адресов, каждый раз, когда встречается адрес, надо значение изменять на число value
// map[адрес][сумма]

// изменение по абсолютной величине - это просто сумма всех именений без учета знака?
// т.е. здесь: +4, -3, -2 абсолютное изменение это 9 или -1?

var addresses = make(map[uint32]string)

func main(){

	last := getSampleBlock()

	balances := make(map[uint32]*uint256.Int)


	for _, trans := range last.Transactions {

		var bufValue *uint256.Int
		bufValue, _ = uint256.FromHex("0x0")

		curFromAddr := hash(trans.From)
		curToAddr := hash(trans.To)
		curValue, err := uint256.FromHex(trans.Value)
		if err != nil {
			panic(err)
		}

		if curValue.IsZero() {
			break
		}

		// should addresses be added before or after zero check?
		if _, ok := addresses[curFromAddr]; !ok {
			addresses[curFromAddr] = trans.From
		}
		if _, ok := addresses[curToAddr]; !ok {
			addresses[curToAddr] = trans.To
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

	fmt.Println(addresses)
	fmt.Println()
	for i,j := range balances {
		fmt.Println(i, ":  ", j)
	}


	//lastBlockNumber, _ := strconv.ParseInt(last, 0, 64)
	//fmt.Println(lastBlockNumber)

	//balances := make(map[string]int)

	//addr, ok := balances[address]

	//for i := lastBlockNumber - 99; i <= lastBlockNumber; i++ {
		//get each block and count whatever is needed to be counted
		//getBlock() // i to hex

}


func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

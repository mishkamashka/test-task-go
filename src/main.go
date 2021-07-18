package main

import (
	"fmt"
	"hash/fnv"

	"github.com/holiman/uint256"
)

// Напиши программу, которая выдаст адрес,
// баланс которого изменился больше остальных
// (по абсолютной величине) за последние 100 блоков.


// map адресов, каждый раз, когда встречается адрес, надо значение изменять на число value
// map[адрес][сумма]

// что-то делать с value uint256...
// может стоит сразу отметать нули

var addresses = make(map[uint32]string)

func main(){
	//last := getSampleBlock()

	//balances := make(map[uint64]*uint256.Int)


	var a,b,c,d *uint256.Int
	a, _ = uint256.FromHex("0xfffffffffffffffffffffffffffffffffffffffffffffffc8fbb078f9332b04a")
	b, _ = uint256.FromHex("0xde4289155b36ecc")
	c, _ = uint256.FromHex("0xaa9f075c200000")
	d, _ = uint256.FromHex("0xaa9f075c200000")

	c.Add(a,d.Neg(b))

	fmt.Print(a, " + ", b, " = ", c)

	//for _, trans := range last.Transactions {
	//
	//	curFromAddr := hash(trans.From)
	//	curToAddr := hash(trans.To)
	//
	//	if _, ok := addresses[curFromAddr]; !ok {
	//		addresses[curFromAddr] = trans.From
	//	}
	//	if _, ok := addresses[curToAddr]; !ok {
	//		addresses[curToAddr] = trans.To
	//	}
	//
	//	transValue, _ := strconv.ParseUint(trans.Value, 0, 64)
	//	fmt.Println(transValue)
	//
	//	if transValue == 0 {
	//		break
	//	}

		//if value, ok := balances[curFromAddr]; ok {
		//	fmt.Println(value)
		//
		//	balances[curFromAddr] = value - trans.Value
		//} else {
		//	fmt.Println("no")
		//	balances[curFromAddr] = - trans.Value
		//}
		//
		//if value, ok := balances[curToAddr]; ok {
		//	fmt.Println(value)
		//	balances[curToAddr] = value + trans.Value
		//} else {
		//	fmt.Println("no")
		//	balances[curToAddr] = trans.Value
		//}
	//}







	//lastBlockNumber, _ := strconv.ParseInt(last, 0, 64)
	//fmt.Println(lastBlockNumber)

	//balances := make(map[string]int)

	//addr, ok := balances[address]

	//for i := lastBlockNumber - 99; i <= lastBlockNumber; i++ {
		//get each block and count whatever is needed to be counted
		//getBlock() // i to hex

	//}

}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

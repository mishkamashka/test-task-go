package main

import "github.com/holiman/uint256"

type Counter struct {
	value		*uint256.Int
	lastSign	bool			//true +, false -
}

func NewCounter(v uint256.Int, sign bool) *Counter {
	return &Counter{
		value:    &v,
		lastSign: sign,
	}
}

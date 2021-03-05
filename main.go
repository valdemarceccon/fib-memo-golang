package main

import (
	"fmt"
	"math/big"
)

func main() {
	for i := 1; i <= 1000; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fibFast(uint(i)))
	}
}

func fibSlow(n uint) *big.Int {
	a := big.NewInt(1)
	if n <= 2 {
		return a
	}
	result := big.Int{}
	result.Add(fibSlow(n-1), fibSlow(n-2))
	return &result
}

func fibFast(n uint) *big.Int {
	m := make(map[uint]*big.Int)
	return fibFastInternal(n, m)
}

func fibFastInternal(n uint, m map[uint]*big.Int) *big.Int {
	if n <= 2 {
		return big.NewInt(1)
	}
	if v, ok := m[n]; ok {
		return v
	}
	ret := &big.Int{}
	ret.Add(fibFastInternal(n-1, m), fibFastInternal(n-2, m))
	m[n] = ret
	return ret
}

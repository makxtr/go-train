package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(uniqN(10))

}

func uniqN(n int) []int64 {
	res := make([]int64, n)
	ints := make(map[int64]struct{}, n)
	var a *big.Int

	for i, _ := range res {
		for {
			a, _ = rand.Int(rand.Reader, big.NewInt(100))
			if _, ok := ints[a.Int64()]; !ok {
				ints[a.Int64()] = struct{}{}
				break
			}
		}

		res[i] = a.Int64()
	}

	return res
}

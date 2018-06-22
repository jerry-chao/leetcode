package main

import (
	"fmt"
)

// main
func main() {
	fmt.Println("hello")
	fmt.Println(numJewelsInStores("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStores("z", "ZZ"))
	fmt.Println(makeStoresHash("ZZ"))
}

func numJewelsInStones(J string, S string) int {
	sHash := makeStoresHash(S)
	var sum int
	for _, value := range J {
		sum = sum + sHash[value]
	}
	return sum
}

func makeStoresHash(S string) map[rune]int {
	ret := make(map[rune]int)
	for _, value := range S {
		ret[value] = ret[value] + 1
	}
	return ret
}

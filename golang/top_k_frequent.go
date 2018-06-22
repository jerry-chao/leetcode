package main

import (
	"fmt"
	"sort"
)

type myInt struct {
	frequent int
	value    int
}

type myIntList []myInt

func (m myIntList) Len() int {
	return len(m)
}

func (m myIntList) Less(i, j int) bool {
	return m[i].frequent < m[j].frequent
}

func (m myIntList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	fmt.Println("hello")
	fmt.Println(calcFrequent([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	frequents := calcFrequent(nums)
	var l myIntList
	for key, value := range frequents {
		fmt.Println("xxxx", key, value, l)
		l = add(l, myInt{value, key}, k)
	}
	return getkeys(l)
}

func getkeys(myNums myIntList) []int {
	var ret []int
	for _, value := range myNums {
		ret = append(ret, value.value)
	}
	return ret
}

func add(l myIntList, new myInt, k int) myIntList {
	if len(l) < k {
		l = append(l, new)
	} else {
		if l[0].frequent < new.frequent {
			l[0] = new
			sort.Sort(l)
			return l
		}

	}
	sort.Sort(l)
	return l
}

func calcFrequent(nums []int) map[int]int {
	ret := make(map[int]int)
	for _, value := range nums {
		ret[value] = ret[value] + 1
	}
	return ret
}

package main

import (
	"log"
	"sort"
)

type myInt struct {
	value int
	index int
}

type myIntList []myInt

func (m myIntList) Len() int {
	return len(m)
}

func (m myIntList) Less(i, j int) bool {
	return m[i].value < m[j].value
}

func (m myIntList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	log.Println("Hello world")
	// a := []int{1, 2, 3, 4, 7, 6}
	// target := 5
	// sumIndexList := twoSum(a, target)

	// log.Println(sumIndexList)
	// log.Println(twoSum([]int{}, 5))
	// log.Println(twoSum([]int{2, 7, 11, 15}, 9))
	// log.Println(twoSum([]int{3, 2, 4}, 6))
	// log.Println(twoSum1([]int{3, 2, 4}, 6))
	log.Println(twoSumHash([]int{2, 7, 11, 15}, 9))
	log.Println(twoSumHash([]int{3, 2, 3}, 6))
	log.Println(twoSumSort([]int{2, 7, 11, 15}, 9))
	log.Println(twoSumSort([]int{3, 2, 3}, 6))
}

func twoSumSort(nums []int, target int) []int {
	var myNums myIntList
	for index, value := range nums {
		myNums = append(myNums, myInt{value, index})
	}
	sort.Sort(myNums)
	listlen := len(myNums)
	var sumIndexList []int
	i, j := 0, listlen-1
	for i < listlen && i < j {
		if myNums[i].value+myNums[j].value == target {
			sumIndexList = append(sumIndexList, myNums[i].index)
			sumIndexList = append(sumIndexList, myNums[j].index)
			i++
			j--
		} else if myNums[i].value+myNums[j].value < target {
			i++
		} else {
			j--
		}
	}
	return sumIndexList
}

func twoSumForce(nums []int, target int) []int {
	sumlen := len(nums)
	var ret []int
	for i := 0; i < sumlen; i++ {
		for j := i + 1; j < sumlen; j++ {
			if nums[i]+nums[j] == target {
				ret = append(ret, i, j)
			}
		}
	}
	return ret
}

func twoSumHash(nums []int, target int) []int {
	mapnums := make(map[int]int)
	var compent int
	for index, value := range nums {
		compent = target - value
		if compentIndex, ok := mapnums[compent]; ok {
			return []int{compentIndex, index}
		}
		mapnums[value] = index
	}
	return []int{}
}

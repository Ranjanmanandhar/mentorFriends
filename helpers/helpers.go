package helpers

import (
	"fmt"
	"io/ioutil"
)

func RemoveDuplicates(arr []int) []int {
	uniqueMap := make(map[int]bool)
	uniqueSlice := []int{}

	for _, num := range arr {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqueSlice = append(uniqueSlice, num)
		}
	}

	return uniqueSlice
}

func WriteToFile(name string, data string) {
	err := ioutil.WriteFile(name, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func UpdateArray(arr []int, old int, new int) []int {
	index := -1
	for i, num := range arr {
		if num == old {
			index = i
			break
		}
	}

	if index != -1 {
		arr = append(arr[:index], arr[index+1:]...)
		arr = append(arr, new)
	} else {
		fmt.Println("Number", old, "not found in the array.")
	}
	return arr
}

func DeleteFromArray(arr []int, del int) []int {
	index := -1
	for i, num := range arr {
		if num == del {
			index = i
			break
		}
	}

	if index != -1 {
		arr = append(arr[:index], arr[index+1:]...)
	} else {
		fmt.Println("Number", del, "not found in the array.")
	}
	return arr
}

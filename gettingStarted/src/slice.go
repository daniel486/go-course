package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var unorderedSlice []int = make([]int, 0)
	var scannedString string
	for strings.ToLower(scannedString) != "x" {
		fmt.Printf("Previous slice state: %v\n", unorderedSlice)
		fmt.Printf("Please type an int to sort it or x to exit:")
		fmt.Scan(&scannedString)
		convertedInt, err := strconv.Atoi(scannedString)
		if err == nil {
			unorderedSlice = append(unorderedSlice, convertedInt)
			sort.SliceStable(unorderedSlice, func(i, j int) bool {
				return unorderedSlice[i] < unorderedSlice[j]
			})
		}
	}
}

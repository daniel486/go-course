package main

import (
	"fmt"
)

func main() {
	var floatToTrunc float64 = -1
	var truncInt int
	fmt.Printf("Please type a decimal number to truncate it or 0 to exit:")
	for floatToTrunc != 0 {
		fmt.Scan(&floatToTrunc)
		truncInt = int(floatToTrunc)
		fmt.Printf("The truncated value is: %d\n", truncInt)
		fmt.Printf("Please type another decimal or 0 to exit:")
	}

}

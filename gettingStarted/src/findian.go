package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var scannedString string
	fmt.Printf("Please input a text: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		stringAux := scanner.Text()
		scannedString = stringAux
	}
	scannedString = strings.ToLower(scannedString)
	if strings.HasPrefix(scannedString, "i") &&
		strings.Contains(scannedString, "a") &&
		strings.HasSuffix(scannedString, "n") {

		fmt.Printf("Found!")
	} else {

		fmt.Printf("Not Found!")
	}
}

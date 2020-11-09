package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var personData map[string]string
	personData = make(map[string]string)
	fmt.Printf("Please input the name of the person: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		stringAux := scanner.Text()
		personData["name"] = stringAux
	}
	fmt.Printf("Please input the address of the person: ")
	if scanner.Scan() {
		stringAux := scanner.Text()
		personData["address"] = stringAux
	}
	jsonObject, err := json.Marshal(personData)
	if(err == nil){
		jsonString := string(jsonObject)
		fmt.Println(jsonString)
	}
}
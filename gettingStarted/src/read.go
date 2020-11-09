package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	nameSlice := make([]name, 0)
	fmt.Printf("Please type the name of the file: ")
	fmt.Scan(&fileName)
	fileData, err := ioutil.ReadFile(fileName)
	if(err == nil){
		namesArr := strings.Split(string(fileData), "\n")
		for i := range namesArr {
			auxArr := strings.Split(namesArr[i], " ")
			auxName := new(name)
			auxName.fname = auxArr[0]
			auxName.lname = auxArr[1]
			nameSlice = append(nameSlice, *auxName)
		}
	}

	for i := range nameSlice {
		fmt.Println(nameSlice[i].fname, nameSlice[i].lname)
	}
}

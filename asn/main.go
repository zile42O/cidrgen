package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type AsnData struct {
	AsnData []AsnStruct `json:"matches"`
}

type AsnStruct struct {
	AsnName string `json:"asn"`
}

func main() {
	jsonFile, err := os.Open("./file.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println("\n\x1b[36m [ \x1b[31m* \x1b[36m]\x1b[0m Successfuly Opened > \x1b[36m [ \x1b[93mfile.json \x1b[36m] \x1b[0m")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var get AsnData
	json.Unmarshal(byteValue, &get)

	//Writing to asn_list.txt
	f, err := os.OpenFile("asn_list.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("\n\x1b[36m [ \x1b[31m* \x1b[36m]\x1b[0m Writing ASN list to > \x1b[36m [ \x1b[93masn_list.txt \x1b[36m] \x1b[0m")
	for i := 0; i < len(get.AsnData); i++ {
		if _, err := f.WriteString(get.AsnData[i].AsnName + " \n"); err != nil {
			log.Println(err)
		}
	}
	f.Close()
	fmt.Println("\n\x1b[36m [ \x1b[31m* \x1b[36m]\x1b[0m Checking & Fixing duplicates in > \x1b[36m [ \x1b[93masn_list.txt \x1b[36m] \x1b[0m")
	// read the lines
	line, _ := ioutil.ReadFile("asn_list.txt")
	// turn the byte slice into string format
	strLine := string(line)
	// split the lines by a space, can also change this
	lines := strings.Split(strLine, " ")
	// remove the duplicates from lines slice (from func we created)
	RemoveDuplicates(&lines)
	// get the actual file
	f, err = os.OpenFile("asn_list.txt", os.O_APPEND|os.O_WRONLY, 0600)
	// err check
	if err != nil {
		log.Println(err)
	}
	// delete old one
	os.Remove("asn_list.txt")
	// create it again
	os.Create("asn_list.txt")
	// go through your lines
	for e := range lines {
		// write to the file without the duplicates
		f.Write([]byte(lines[e] + " ")) // added a space here, but you can change this
	}
	// close file
	f.Close()
}

func RemoveDuplicates(lines *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *lines {
		if !found[x] {
			found[x] = true
			(*lines)[j] = (*lines)[i]
			j++
		}
	}
	*lines = (*lines)[:j]
}

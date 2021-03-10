package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {
	API_KEY := "https://api.bgpview.io/asn/%s/prefixes"

	var err error
	fmt.Println("\n\x1b[36m [ \x1b[31m* \x1b[36m]\x1b[0m Getting CIDR Ranges from > \x1b[36m [ \x1b[93m asn/asn_list.txt \x1b[36m] \x1b[0m")

	file, err := os.Open("../asn/asn_list.txt")
	if err != nil {
		fmt.Println("\n\x1b[36m [ \x1b[31m* \x1b[36m]\x1b[0m FILE `asn_list.txt` > \x1b[36m [ \x1b[93m Not exists \x1b[36m] \x1b[0m")
		//log.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	f, err := os.OpenFile("cidr_list.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for scanner.Scan() {
		asnnumber := scanner.Text()
		fmt.Println("\n\x1b[36m [ \x1b[31m* \x1b[36m]\x1b[0m Scaning > \x1b[36m [ \x1b[93m " + asnnumber + " \x1b[36m] \x1b[0m")
		asnnumber = strings.ReplaceAll(asnnumber, "AS", "")
		// request http api
		url := fmt.Sprintf(API_KEY, asnnumber) //prepare url

		res, err := http.Get(url)
		if err != nil {
			fmt.Println("\n\x1b[36m [ \x1b[31m* \x1b[36m]\x1b[0m API KEY > \x1b[36m [ \x1b[93m Not Working \x1b[36m] \x1b[0m")
			return
		}
		//fmt.Println("\n\x1b[36m [ \x1b[31m* \x1b[36m]\x1b[0m API KEY > \x1b[36m [ \x1b[93m Working \x1b[36m] \x1b[0m")

		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		newasn := gjson.Get(string(body), "data.ipv4_prefixes.#.prefix")
		for _, name := range newasn.Array() {
			//println(name.String())
			if _, err := f.WriteString("\n" + name.String() + " "); err != nil {
				log.Println(err)
			}
		}
	}
}

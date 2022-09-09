package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SubDomain struct {
	IssuerCaID     int    `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	ID             int64  `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
	SerialNumber   string `json:"serial_number"`
}

func main() {
	var site_name = os.Args[1]
	resp, err := http.Get("https://crt.sh/?q=" + site_name + "&output=json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var temp []SubDomain

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	marshal_err := json.Unmarshal(body, &temp)
	// fmt.Println(marshal_err)
	if marshal_err != nil {
		fmt.Println(marshal_err)
	}
	for i := range temp {
		fmt.Println(temp[i].CommonName)
		fmt.Println(temp[i].NameValue)
	}

	// log.Println(string(body))
}

package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"godata/edmx"
)

//const url = "https://httpbin.org/ip"
const url = "https://services.odata.org/V2/OData/OData.svc/"

func fetchMetadata(url string) {
	response, err := http.Get(url + "$metadata")

	if err != nil {
		fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("The HTTP request body read failed with error %s\n", err)
	}

	fmt.Println(string(data))

	var metadata edmx.Edmx

	err = xml.Unmarshal(data, &metadata)
	if err != nil {
		fmt.Errorf("The metadata parsing failed with error %s\n", err)
	}

	metadataJson, err := json.MarshalIndent(metadata, "", "  ")

	fmt.Printf("%s\n", metadataJson)
}

func fetchUser() {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func main() {
	fmt.Println("Starting the application...")

	fetchMetadata(url);

	fmt.Println("Terminating the application...")
}

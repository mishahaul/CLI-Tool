package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Product struct {		//initializing a struct which will hold products
	ID int `json:"ID"`
	Name string `json:"Name"`
	Description string `json:"Description"`
	Price int `json:"Price"`
	SalesPrice int `json:"Sales Price"`
	Features []string `json:"Features"`
}

func main(){
	content, err := ioutil.ReadFile("stuff.json") // read json file
	if err != nil {
		fmt.Println(err.Error())
	}

	var products []Product
	err2 := json.Unmarshal(content, &products) // unmarshal decoded json into products
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	// for _, v := range products {
	// 	fmt.Printf("%s\n", v.Name)
	// }
	// fmt.Println(products)
}

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

var products []Product

func MainMenu() int {
	var option int
	fmt.Println("Main menu:")
	fmt.Println("1. Show Products")
	fmt.Println("2. Filter Products")
	fmt.Printf("Your choice is: \n")
	fmt.Scanf("%d", &option)
	return option
}

func ShowProductsMenu(p []Product) int {
	var option int
	fmt.Println("Show Products Menu:")
	// fmt.Println(p[0])
	fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\n%s\n", p[0].ID, p[0].Name, 
				p[0].Description, p[0].Price, p[0].SalesPrice, p[0].Features)
	fmt.Println("1. Edit product")
	fmt.Println("2. Next")
	fmt.Println("3. Previous")
	fmt.Println("4. Back to menu")
	fmt.Printf("Your choice is: \n")
	fmt.Scanf("%d", &option)
	return option
}


func main(){
	content, err := ioutil.ReadFile("stuff.json") // read json file
	if err != nil {
		fmt.Println(err.Error())
	}

	
	err2 := json.Unmarshal(content, &products) // unmarshal decoded json into products
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	// MainMenu()
	var option1 int = MainMenu()
	var option2 int
	switch option1 {
	case 1: 
		fmt.Println("Show Products Menu: ")
		option2 = ShowProductsMenu(products)
	case 2:  
		fmt.Println("Filter Product Menu: ")
	default:
		fmt.Println(`Enter "1" to Show Products or press "2" to Filter them`)
		MainMenu()
	}
	// fmt.Println(option1)
	fmt.Println(option2)
	switch option2 {
	case 1: 
		fmt.Println("1. Edit product ")
	case 2:  
		fmt.Println("2. Next")
	case 3:  
		fmt.Println("3. Previous")
	case 4:  
		fmt.Println("4. Back to menu")
	default:
		fmt.Println("Enter one of available options!")
		ShowProductsMenu(products)
	}

}

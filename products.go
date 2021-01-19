package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
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

func MainMenu() {
	for {
		var option int
		fmt.Println("Main menu:")
		fmt.Println("1. Show Products")
		fmt.Println("2. Filter Products")
		fmt.Println("3. Exit")
		fmt.Printf("Your choice is: ")
		n, err := fmt.Scan(&option)
		fmt.Println(n, err)
		fmt.Println(option)
		fmt.Printf("%T\n", option)

		if err != nil || option < 1 || option > 3 {
			fmt.Println(`Enter "1" to Show Products, press "2" to Filter them or "3" to Exit`)
			// MainMenu()
		}
	
		switch option {
			case 1: 
				fmt.Println("Show Products Menu: ")
				ShowProductsMenu(products)	
			case 2:  
				fmt.Println("Filter Product Menu: ")
				FilterProductMenu(products)
			case 3:  
				os.Exit(0)
			// default:
			// 	fmt.Println(`Enter "1" to Show Products, 
			// 	press "2" to Filter them or "3" to Exit`)
		}
	}
}

func ShowProductsMenu(p []Product) {
	i := 0

	for {
		var option int
		// fmt.Println("Show Products Menu:")
		// fmt.Println(p[0])
		fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\n%s\n", p[i].ID, p[i].Name, 
					p[i].Description, p[i].Price, p[i].SalesPrice, p[i].Features)
		fmt.Println("1. Edit product")
		fmt.Println("2. Next")
		fmt.Println("3. Previous")
		fmt.Println("4. Back to menu")
		fmt.Printf("Your choice is: \n")
		fmt.Scan(&option)
		// fmt.Println("test")
		switch option {
			case 1: 
				// fmt.Println("Edit product")	
				EditProductMenu(&p[i])
			case 2:  
				fmt.Println("Next")
				i++
				if i == len(p) {
					i = 0
				}
			case 3:  
				fmt.Println("Previous")
				i--
				if i == -1 {
					i = len(p) - 1
				}
			case 4:  
				MainMenu()
			default:
				fmt.Println("Press one of available options!")
		}
	}
}

func EditProductMenu(p *Product) {
	for {
		var option int
		fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\n%s\n", p.ID, p.Name, 
					p.Description, p.Price, p.SalesPrice, p.Features)
		fmt.Println("Edit Product Menu:")
		fmt.Println("1. Name")
		fmt.Println("2. Description")
		fmt.Println("3. Price")
		fmt.Println("4. Sales Price")
		fmt.Println("5. Back to menu")
		fmt.Printf("Your choice is: \n")
		fmt.Scan(&option)
		switch option {
			case 1: 
				fmt.Println("What is new Name gonna be:")	
				fmt.Scan(&p.Name)
			case 2:  
				fmt.Println("What is new Description gonna be:")
				fmt.Scan(&p.Description)
			case 3:  
				fmt.Println("Updated Price is:")
				fmt.Scan(&p.Price)
			case 4:  
				fmt.Println("Is it on Sale? Final Price is:")
				fmt.Scan(&p.SalesPrice)
			case 5:  
				// MainMenu()
				ShowProductsMenu(products)
			default:
				fmt.Println("Press one of available options!")
		}
	}
}

func FilterProductMenu(p []Product) {
	i := 0
	for {
		var option int
		fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\n%s\n", p[i].ID, p[i].Name, 
		p[i].Description, p[i].Price, p[i].SalesPrice, p[i].Features)
		fmt.Println("Choose field to filter by:")
		fmt.Println("1. Name")
		fmt.Println("2. Description")
		fmt.Println("3. Price")
		fmt.Println("4. Sales Price")
		fmt.Println("5. Features")
		fmt.Println("6. Back to menu")
		fmt.Printf("Your choice is: ")
		fmt.Scan(&option)
		switch option {
			case 1: 
				var userInput string
				fmt.Println("Which Name to filter by?")
				fmt.Scan(&userInput)
				if strings.Compare(userInput, p[i].Name) == 0 {
					fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\n%s\n", p[i].ID, p[i].Name, 
					p[i].Description, p[i].Price, p[i].SalesPrice, p[i].Features)
				}
			case 2:  
				fmt.Println("Description")
			case 3:  
				fmt.Println("Price")
				
			case 4:  
				fmt.Println("Sales Price")
			case 5:  
				fmt.Println("Features")
			case 6:  
				MainMenu()
			default:
				fmt.Println("Press one of available options!")
		}
		i++
	}
}

func main() {
	content, err := ioutil.ReadFile("stuff.json") // read json file
	
	if err != nil {
		fmt.Println(err.Error())
	}

	err2 := json.Unmarshal(content, &products) // unmarshal decoded json into products
	
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	MainMenu()

}

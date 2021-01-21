package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"strconv"
	// "errors"
	// "strings"
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
		// fmt.Scan(&option)
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(n, err)
		// fmt.Println(option)
		// fmt.Printf("%T\n", option)

		if option < 1 || option > 3 {
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
	
	for {
		var option int
		fmt.Println("Choose field to filter by:")
		fmt.Println("1. Name")
		fmt.Println("2. Price")
		// fmt.Println("3. Description")
		// fmt.Println("4. Sales Price")
		// fmt.Println("5. Features")
		fmt.Println("3. Back to menu")
		fmt.Printf("Your choice is: ")
		fmt.Scan(&option)
		switch option {
			case 1: 
				var userInput string
				fmt.Println("Which Name to filter by?")
				fmt.Scan(&userInput)
				for i, val := range p {
					if userInput == val.Name {
						fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\n%s\n", p[i].ID, p[i].Name, 
						p[i].Description, p[i].Price, p[i].SalesPrice, p[i].Features)
					} else if userInput != val.Name && i == len(p) - 1 {
						fmt.Printf("We don't have any item matching %v\n", userInput)
					}
				}
			case 2:  
				fmt.Println("Price")
				filterByPrice(p)
				// FilterProductMenu(products)
			// case 3:  
			// 	fmt.Println("Description")
				
			// case 4:  
			// 	fmt.Println("Sales Price")
			// case 5:  
			// 	fmt.Println("Features")
			case 3:  
				MainMenu()
			default:
				fmt.Println("Press one of available options!")
		}
	}
}

func filterByPrice(p []Product) {
	// var err1, err2 error
	// err1 := errors.New("Flag UP")
	// err2 := errors.New("Flag UP")
	// fmt.Println(err1, err2)

	for {
		var minPrice string
		var maxPrice string
		fmt.Println("What is the minimum item Price to filter by?")
		fmt.Scan(&minPrice)
		min, err1 := strconv.Atoi(minPrice)
		// fmt.Println(err1)
		if err1 != nil {
			fmt.Println("Nunbers only")
			// fmt.Println("Please enter minimum price to look up from")
			// fmt.Scan(&minPrice)
			continue
		}
		// fmt.Printf("%d\n", min)
		// fmt.Printf("%T\n", min)

		fmt.Println("What is the maximum item Price to filter by?")
		fmt.Scan(&maxPrice)
		max, err2 := strconv.Atoi(maxPrice)
		// fmt.Println(err2)
		if err2 != nil {
			fmt.Println("Nunbers only")
			// fmt.Println("Please enter maximum price to look up from")
			continue
		}
		// fmt.Printf("%d\n", max)
		// fmt.Printf("%T\n", max)

		for i, val := range p {
			if min <= val.Price && val.Price <= max {
				fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\n%s\n", p[i].ID, p[i].Name, 
				p[i].Description, p[i].Price, p[i].SalesPrice, p[i].Features)
			} 
		}
		if err1 == nil && err2 == nil {
			break
		}
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

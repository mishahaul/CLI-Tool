package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type Product struct { 						//initializing a struct which will hold products
	ID          int      `json:"ID"`
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Price       int      `json:"Price"`
	SalesPrice  int      `json:"Sales Price"`
	Features    []string `json:"Features"`
}

var products []Product

func MainMenu() {
	cmdClear()
	for {
		var option int
		printMainMenu()
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
		if option < 1 || option > 3 {
			fmt.Println(`Enter "1" to Show Products, press "2" to Filter them or "3" to Exit`)
		}
		switch option {
		case 1:
			ShowProductsMenu(products)
		case 2:
			FilterProductMenu(products)
			fmt.Println("Filter Product Menu: ")
		case 3:
			os.Exit(0)
		}
	}
}

func printMainMenu() {
	fmt.Println("Main menu:")
	fmt.Println("1. Show Products")
	fmt.Println("2. Filter Products")
	fmt.Println("3. Exit")
	fmt.Printf("Your choice is: ")
}

func ShowProductsMenu(p []Product) {
	cmdClear()
	i := 0

	for {
		var option int
		fmt.Println("Product: ")
		fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\nFeatures is %s\n\n", p[i].ID, p[i].Name,
			p[i].Description, p[i].Price, p[i].SalesPrice, p[i].Features)
		printSPMenu()
	
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
		switch option {
		case 1:
			EditProductMenu(&p[i])
		case 2:
			cmdClear()
			fmt.Println("Next")
			moveIndex(&i, 1, len(p))
		case 3:
			cmdClear()
			fmt.Println("Previous")
			moveIndex(&i, -1, len(p))
		case 4:
			MainMenu()
		default:
			fmt.Println("Press one of available options!")
		}
	}
}

func printSPMenu() {
	fmt.Println("1. Edit product")
	fmt.Println("2. Next")
	fmt.Println("3. Previous")
	fmt.Println("4. Back to menu")
	fmt.Printf("Your choice is: ")
}

func moveIndex(i *int, n, l int) {
	*i += n
	if *i == l {
		*i = 0
	}
	if *i < 0 {
		*i = l - 1
	}
}

func EditProductMenu(p *Product) {
	cmdClear()
	for {
		var option int
		fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\nFeatures is %s\n\n", p.ID, p.Name,
			p.Description, p.Price, p.SalesPrice, p.Features)
		printEPMenu()
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
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
			ShowProductsMenu(products)
		default:
			fmt.Println("Press one of available options!")
		}
	}
}

func printEPMenu() {
	fmt.Println("Edit Product by:")
	fmt.Println("1. Name")
	fmt.Println("2. Description")
	fmt.Println("3. Price")
	fmt.Println("4. Sales Price")
	fmt.Println("5. Back to menu")
	fmt.Printf("Your choice is: ")
}

func FilterProductMenu(p []Product) {
	cmdClear()
	for {
		var option int
		printFPMenu()
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
		switch option {
		case 1:
			cmdClear()
			fmt.Println("NAme")
			filterByName(p)
		case 2:
			cmdClear()
			fmt.Println("Price")
			filterByPrice(p)
		case 3:
			MainMenu()
		default:
			fmt.Println("Press one of available options!")
		}
	}
}
func printFPMenu() {
	fmt.Println("Choose field to filter by:")
	fmt.Println("1. Name")
	fmt.Println("2. Price")
	fmt.Println("3. Back to menu")
	fmt.Printf("Your choice is: ")
}

func filterByName(p []Product) {
	var userInput string
	fmt.Println("Which Name to filter by?")
	fmt.Scan(&userInput)

	for i, val := range p {
		if userInput == val.Name {
			fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\nFeatures is %s\n\n", p[i].ID, p[i].Name,
			p[i].Description, p[i].Price, p[i].SalesPrice, p[i].Features)
		} else if userInput != val.Name && i == len(p)-1 {
			fmt.Printf("We don't have any item matching %v\n", userInput)
		}
	}
}

func filterByPrice(p []Product) {
	for {
		var minPrice string
		var maxPrice string
		fmt.Println("What is the minimum item Price to filter by?")
		fmt.Scan(&minPrice)
		min, err := strconv.Atoi(minPrice)
		if err != nil {
			fmt.Println("Nunbers only")
			continue
		}

		fmt.Println("What is the maximum item Price to filter by?")
		fmt.Scan(&maxPrice)
		max, err := strconv.Atoi(maxPrice)
		if err != nil {
			fmt.Println("Nunbers only")
			continue
		}

		for i, val := range p {
			if min <= val.Price && val.Price <= max {
				fmt.Printf("ID-%d\n%s\n%s\nPrice: %d\nSales Price: %d\n%s\n", p[i].ID, p[i].Name,
					p[i].Description, p[i].Price, p[i].SalesPrice, p[i].Features)
			}
		}
		return
	}
}

func cmdClear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	content, err := ioutil.ReadFile("stuff.json") // read json file
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	err = json.Unmarshal(content, &products) // unmarshal decoded json into products
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	MainMenu()
}

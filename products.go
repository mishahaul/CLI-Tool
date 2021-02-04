package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"./dbConnection"

	_ "github.com/lib/pq"
)

type Product struct { //initializing a struct which will hold products
	ID          int      `json:"ID"`
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Price       int      `json:"Price"`
	SalesPrice  int      `json:"SalesPrice"`
	Features    []string `json:"Features"`
}

var products []Product

var db *sql.DB

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
		var option string
		printFPMenu()
		_, err := fmt.Scan(&option)
		intOpt, err := strconv.Atoi(option)
		if err != nil {
			fmt.Println("Nunbers only")
			continue
		}
		switch intOpt {
		case 1:
			cmdClear()
			fmt.Println("Name")
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
		if max < min {
			fmt.Println("Maximum item price have to be greater, than minimum price")
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

func insertProduct(p []Product) error {
	stmt, err := db.Prepare("TRUNCATE TABLE products CASCADE")
	if err != nil {
		log.Fatalf("truncate failed: %v", err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		log.Fatalf("truncate statment failed: %v", err)
	}

	stmt, err = db.Prepare("INSERT INTO products (product_id, name, description, price, sales_price) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatalf("Prepared statement failed: %v", err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, value := range p {
		if _, err := stmt.Exec(value.ID, value.Name, value.Description, value.Price, value.SalesPrice); err != nil {
			log.Fatalf("Executing statment failed: %v", err)
		}
		// fmt.Println("Your shit uploaded correctly!!!")
	}
	return err
}

func insertFeature(featureMap map[string]int) error {
	stmt, err := db.Prepare("TRUNCATE TABLE features CASCADE")
	if err != nil {
		log.Fatalf("truncate failed: %v", err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		log.Fatalf("truncate statment failed: %v", err)
	}

	stmt, err = db.Prepare("INSERT INTO features (value, feature_id) VALUES ($1, $2)")
	if err != nil {
		log.Fatalf("Prepared statement failed: %v", err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for value, id := range featureMap {
		if _, err := stmt.Exec(value, id); err != nil {
			log.Fatalf("Executing statment failed: %v", err)
		}
		// fmt.Println("Your shit uploaded correctly!!!")
	}
	return err
}

func insertPF(p []Product, featureMap map[string]int) error {
	stmt, err := db.Prepare("TRUNCATE TABLE product_feature")
	if err != nil {
		log.Fatalf("truncate failed: %v", err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		log.Fatalf("truncate statment failed: %v", err)
	}

	stmt, err = db.Prepare("INSERT INTO product_feature (product_id, feature_id) VALUES ($1, $2)")
	if err != nil {
		log.Fatalf("Prepared statement failed: %v", err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, val := range p {
		for _, feature := range val.Features {
			for value, id := range featureMap {
				if feature == value {
					if _, err := stmt.Exec(val.ID, id); err != nil {
						log.Fatalf("Executing statment failed: %v", err)
					}
				}
			}
		// fmt.Println("Your shit uploaded correctly!!!")
		}
	}
	return err
}

func uniqueFeature(p []Product) map[string]int {
	f := make(map[string]int)
	id := 1
	for _, ft := range p {
		for _, val := range ft.Features {
			if _, ok := f[val]; ok {
				continue
			} else {
				f[val] = id
				id++
			}
		}
	}
	return f
}

func main() {
	content, err := ioutil.ReadFile("stuff.json") // read json file
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &products) // unmarshal decoded json into products
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(products)

	// MainMenu()

	dbSettings := dbConnection.Settings{
		User: "postgres",
		Pass: "Logvynets1",
		Name: "postgres",
		Host: "localhost",
		Port: "5432",
	}

	db, err = dbConnection.Connect(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	insertProduct(products)
	f := uniqueFeature(products)
	// fmt.Printf("%v\n", f)
	insertFeature(f)
	insertPF(products, f)
	// fmt.Printf("%v", f)

}

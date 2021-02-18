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

var db *sql.DB

type Product struct { //initializing a struct which will hold products
	ID          int      `json:"ID"`
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Price       int      `json:"Price"`
	SalesPrice  int      `json:"SalesPrice"`
	Features    []string `json:"Features"`
}

// var products []Product

func readJSON(filename string) {
	var products []Product
	content, err := ioutil.ReadFile(filename) // read json file
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &products) // unmarshal decoded json into products
	if err != nil {
		log.Fatal(err)
	}
}

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
			ShowProductsMenu()
		case 2:
			FilterProductsMenu()
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

func ShowProductsMenu() {
	cmdClear()

	i := 1
	var cp int
	row := db.QueryRow("SELECT count(id) FROM products")  // max
	err := row.Scan(&cp)
	if err != nil {
		log.Fatalf("scan of maxid failed")
	}

	for {
		pr := new(Product)
		var option int
		fmt.Println("Product: ")
		if err := db.Ping(); err != nil {
			fmt.Println("no ping")
		}
		row := db.QueryRow("SELECT * FROM products WHERE id = $1", i) 
		err = row.Scan(&pr.ID, &pr.Name, &pr.Description, &pr.Price, &pr.SalesPrice)
		if err != nil {
			log.Fatalf("scan from products failed")
		}
		rows, err := db.Query("select value from features f inner join product_feature pf on pf.feature_id = f.id and pf.product_id = $1", i) 
		if err != nil {
			fmt.Println("select features failed")
		}
		defer rows.Close()
		for rows.Next() {
			var feature string
			err = rows.Scan(&feature)
			if err != nil {
				log.Fatalf("scan from features failed")
			}
			pr.Features = append(pr.Features, feature)
		}
		fmt.Printf("ID - %d\nName: %s\nDescription: %s\nPrice: %d\nSales Price: %d\nFeatures: %v\n", pr.ID, pr.Name,
			pr.Description, pr.Price, pr.SalesPrice, pr.Features)

		printSPMenu()

		_, err = fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
		switch option {
		case 1:
			EditProductMenu(pr)
		case 2:
			cmdClear()
			fmt.Println("Next")
			moveIndex(&i, 1, cp)
		case 3:
			cmdClear()
			fmt.Println("Previous")
			moveIndex(&i, -1, cp)
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
	if *i == l+1 {
		*i = 1
	}
	if *i < 1 {
		*i = l
	}
}

func EditProductMenu(pr *Product) {
	cmdClear()
	for {
		var (
			option     int
			innerQuery string
			param      interface{}
		)

		fmt.Printf("ID - %d\nName: %s\nDescription: %s\nPrice: %d\nSales Price: %d\nFeatures: %v\n", pr.ID, pr.Name,
			pr.Description, pr.Price, pr.SalesPrice, pr.Features)
		printEPMenu()
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
		switch option {
		case 1:
			fmt.Println("What is new Name gonna be:")
			fmt.Scan(&pr.Name)
			innerQuery = "name = $2"
			param = pr.Name
		case 2:
			fmt.Println("What is new Description gonna be:")
			fmt.Scan(&pr.Description)
			innerQuery = "description = $2"
			param = pr.Description
		case 3:
			fmt.Println("Updated Price is:")
			fmt.Scan(&pr.Price)
			innerQuery = "price = $2"
			param = pr.Price
		case 4:
			fmt.Println("Is it on Sale? Final Price is:")
			fmt.Scan(&pr.SalesPrice)
			innerQuery = "sales_price = $2"
			param = pr.SalesPrice
		case 5:

			ShowProductsMenu()
		default:
			fmt.Println("Press one of available options!")
		}
		query := "update products set " + innerQuery + " where id = $1"
		if _, err := db.Exec(query, pr.ID, param); err != nil {
			log.Fatalf("insert statment failed: %v", err)
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

func FilterProductsMenu() {
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
			filterByName()
		case 2:
			cmdClear()
			fmt.Println("Price")
			filterByPrice()
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

func filterByName() {
	pr := Product{}
	var userInput string
	fmt.Println("Which Name to filter by?")
	fmt.Scan(&userInput)
	rows, err := db.Query("select * from products where name = $1", userInput)
	if err != nil {
		fmt.Println("select failed")
	}
	defer rows.Close()
	// if err = rows.Err(); err != nil {
	// 	fmt.Printf("We don't have any item matching %v\n", userInput)
	// }
	// if rows == nil {
	// 	fmt.Printf("We don't have any item matching %v\n", userInput)
	// }

	for rows.Next() {
		// fmt.Println("test")
		err = rows.Scan(&pr.ID, &pr.Name, &pr.Description, &pr.Price, &pr.SalesPrice)
		if err != nil {
			log.Fatalf("scan products failed")
		}
		rowsF, err := db.Query("select value from features f inner join product_feature pf on pf.feature_id = f.id and pf.product_id = $1", pr.ID) 
		if err != nil {
			fmt.Println("select features failed")
		}
		defer rowsF.Close()
		for rowsF.Next() {
			var feature string
			err = rowsF.Scan(&feature)
			if err != nil {
				log.Fatalf("scan from features failed")
			}
			pr.Features = append(pr.Features, feature)
		}
		fmt.Printf("ID - %d\nName: %s\nDescription: %s\nPrice: %d\nSales Price: %d\nFeatures: %v\n", pr.ID, pr.Name,
			pr.Description, pr.Price, pr.SalesPrice, pr.Features)
	}
}

func filterByPrice() {
	// pr := Product{}
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

		rows, err := db.Query("select * from products where price between $1 and $2", minPrice, maxPrice)
		if err != nil {
			fmt.Println("select failed")
		}
		defer rows.Close()
		for rows.Next() {
			pr := Product{}
			err := rows.Scan(&pr.ID, &pr.Name, &pr.Description, &pr.Price, &pr.SalesPrice)
			if err != nil {
				log.Fatalf("scan products failed")
			}
			rowsF, err := db.Query("select value from features f inner join product_feature pf on pf.feature_id = f.id and pf.product_id = $1", pr.ID) 
			if err != nil {
				fmt.Println("select features failed")
			}
			defer rowsF.Close()
			for rowsF.Next() {
				var feature string
				err = rowsF.Scan(&feature)
				if err != nil {
					log.Fatalf("scan from features failed")
				}
				pr.Features = append(pr.Features, feature)   
				
			}
			fmt.Printf("ID - %d\nName: %s\nDescription: %s\nPrice: %d\nSales Price: %d\nFeatures: %v\n", pr.ID, pr.Name,
				pr.Description, pr.Price, pr.SalesPrice, pr.Features)	
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

	stmt, err = db.Prepare("INSERT INTO products (id, name, description, price, sales_price) VALUES ($1, $2, $3, $4, $5)")
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

	stmt, err = db.Prepare("INSERT INTO features (value, id) VALUES ($1, $2)")
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
	var err error

	dbSettings := dbConnection.Settings{
		User: "postgres",
		Pass: "Logvynets1",
		Name: "postgres",
		Host: "localhost",
		Port: "25432",
	}

	db, err = dbConnection.Connect(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	MainMenu()

	// insertProduct(products)
	// f := uniqueFeature(products)
	// // fmt.Printf("%v\n", f)
	// insertFeature(f)
	// insertPF(products, f)
	// // fmt.Printf("%v", f)

}

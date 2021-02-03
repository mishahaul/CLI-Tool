package dbConnection

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func Connect(settings Settings) (*sql.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
				settings.Host, settings.Port, settings.User, settings.Pass, settings.Name)

	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatal(err)
	
	}
	log.Printf("Database connection was created: %s\n", sqlInfo)
	return db, err
}

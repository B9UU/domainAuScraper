package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func insertRow(db *sql.DB, tableName string, data ListingType) (int64, error) {
	insertDataSQL := fmt.Sprintf(`
		INSERT OR IGNORE INTO %s (id,
						address,
						price,
						agents,
						images,
						url,
						beds,
						baths,
						parking

						)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, tableName)

	result, err := db.Exec(
		insertDataSQL,
		data.ID,
		data.ListingModel.Address.Street+","+data.ListingModel.Address.Suburb+","+data.ListingModel.Address.State+","+data.ListingModel.Address.Postcode,
		data.ListingModel.Price,
		data.Branding.AgentNames,
		strings.Join(data.ListingModel.Images, ","),
		data.ListingModel.URL,
		data.Features.Beds,
		data.Features.Baths,
		data.Features.Parking)
	if err != nil {
		return 0, err
	}

	lastInsertID, _ := result.LastInsertId()
	return lastInsertID, nil
}
func openDB(tableName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "scraper.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id			INTEGER PRIMARY KEY,
			address 	TEXT,
			price 		TEXT,
			agents		TEXT,
			images		TEXT,
			url			TEXT,
			beds		INTEGER,
			baths		INTEGER,
			parking		INTEGER

		)
	`, tableName)

	_, err = db.Exec(createTableSQL)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
func main() {
	defer Timer(23)()
	db, err := openDB("first")
	if err != nil {
		log.Fatalln("unable to open database ", err)
	}
	reqPerSec := 5
	r := NewRequester(reqPerSec)

	for i := 1; i <= 3; i++ {

		r.SendRequest(db, i)
		fmt.Println("ss")
	}
	r.Wait()
}

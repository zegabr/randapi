package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"randapi.com/entity"
)

type repo struct {
	db *sql.DB
}

func NewMysqlRepository() Repository {
	db, err := sql.Open("mysql", "user:password@tcp(mysql:3306)/database")
	if err != nil {
		log.Fatal(err)
	}
	return &repo{
		db: db,
	}
}

func (r *repo) GetRow() (*entity.Row, error) {
	query := fmt.Sprintf("SELECT * FROM `languages` WHERE `language` = 'go'")
	result, err := r.db.Query(query)
	if err != nil {
		log.Println("Error getting row:", err)
		return nil, err
	}
	var row entity.Row
	if result.Next() {
		err := result.Scan(&row.Language, &row.Phrase)
		if err != nil {
			log.Println("Error scanning table row:", err)
			return nil, err
		}
	}
	if err = result.Err(); err != nil {
		log.Println("Error iterating table result:", err)
		return nil, err
	}
	return &row, nil
}

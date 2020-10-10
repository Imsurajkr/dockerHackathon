package main

import "log"
import "fmt"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

type Book struct {
	Id uint32
	Name string
}

func main() {
	books := GetAllBooks()
	for _, book := range books {
		fmt.Println(book)
	}
}

func acqConn() *sql.DB {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetAllBooks() []Book {
	var books []Book
	db := acqConn()
	defer db.Close()
	rows, err := db.Query("select * from book")
	if err != nil {
		log.Print(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var (
			Id int
			name string
		)
		rows.Scan(&Id, &name)
		books = append(books, Book{Id: uint32(Id), Name: name})
	}
	return books
}


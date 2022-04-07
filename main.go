package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err := sql.Open("postgres", "user=admin dbname=testdb password=admin sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// Create
	{
		cmd := "INSERT INTO Persons (name, age) VALUES ($1, $2)"
		_, err := Db.Exec(cmd, "name2", 30)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Read
	{
		cmd := "SELECT * FROM Persons WHERE age = $1"
		row := Db.QueryRow(cmd, 20)
		var p Person
		err = row.Scan(&p.Name, &p.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("No row")
			} else {
				log.Println(err)
			}
		}
		fmt.Println(p.Name, p.Age)

		cmd = "SELECT * FROM Persons"
		rows, _ := Db.Query(cmd)
		defer rows.Close()
		var pp []Person
		for rows.Next() {
			var p Person
			err := rows.Scan(&p.Name, &p.Age)
			if err != nil {
				if err == sql.ErrNoRows {
					log.Println("No row1")
				} else {
					log.Println(err)
				}
			}
			pp = append(pp, p)
		}
		for _, p := range pp {
			fmt.Println(p.Name, p.Age)
		}
	}

	// Update
	{
		cmd := "UPDATE Persons SET age = $1 WHERE name = $2"
		_, err = Db.Exec(cmd, 35, "name2")
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Delete
	{
		cmd := "DELETE FROM Persons WHERE name = $1"
		_, err = Db.Exec(cmd, "name2")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

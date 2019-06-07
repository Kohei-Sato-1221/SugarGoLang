package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func main(){
	lesson79()
}

var DbConnection *sql.DB

type Person struct {
	Name string
	Age int
}

func lesson79(){
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(name STRING, age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	
//	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
//	_, err = DbConnection.Exec(cmd, "Kohei", 31)
//	_, err = DbConnection.Exec(cmd, "Kanako", 28)
//	if err != nil {
//		log.Fatal(err)
//	}

//	cmd = "UPDATE person SET age =? WHERE name = ?"
//	_, err = DbConnection.Exec(cmd, 100, "Kanako")
//	if err != nil {
//		log.Fatalln(err)
//	}

//	cmd = "SELECT * FROM person"
//	rows, _ := DbConnection.Query(cmd)
//	defer rows.Close()
//	var pp []Person
//	for rows.Next() {
//		var p Person
//		err := rows.Scan(&p.Name, &p.Age)
//		if err != nil {
//			log.Println(err)
//		}
//		pp = append(pp, p)
//	}
//	for _, p := range pp {
//		fmt.Println(p.Name, p.Age)
//	}
	
	cmd = "SELECT * FROM person where age = ?"
	row := DbConnection.QueryRow(cmd, 1000)
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
}







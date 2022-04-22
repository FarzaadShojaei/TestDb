package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	//connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=postgres user=postgres password=")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to Connect :%v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to Database!")
	//test my connection

	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot Ping Database")
	}
	log.Println("Pinged Database!")

	//get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	//insert a row
	query := `insert into users(first_name,last_name) values ($1,$2)`
	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a row")

	//update a row
	stmt := `update users set first_name=$1 where first_name=$2`
	_, err = conn.Exec(stmt, "Jackie", "Jack")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Updated one or more rows")

	//get one row by id
	query = `select id, first_name, last_name from users where id = $1`

	var firstName, lastName string
	var id int

	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow returns", id, firstName, lastName)

	//delete a row
	query = `delete from users where id= $1`
	_, err = conn.Exec(query, 6)
	if err != nil {
		log.Fatal(err)

	}
	log.Println("Deleted a Row")
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select_id,first_name,last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)

		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is ", id, firstName, lastName)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("----------------------------")

	return nil
}

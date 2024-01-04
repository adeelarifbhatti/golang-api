package main

import(
	"log"
	"testing"
	"fmt"
)
var a App

func  TestMain(m *testing.M) {
	a.Start("root",DbPassword,DBName)
	createDB()
	createLanguage("Julie")
	m.Run()
}
func createDB() {
	createDatabase := `CREATE DATABASE testing;`
	useDatabase := `use testing;`
	createTable := `CREATE TABLE testing(id int Primary Key NOT NULL AUTO_INCREMENT,name VARCHAR(255));`
	_, err := a.DB.Exec(createDatabase)
	if err != nil {
		log.Fatal(err)
	}
	_, err = a.DB.Exec(useDatabase)
	if err != nil {
		log.Fatal(err)
	}
	_, err = a.DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
func createLanguage(name string){
	query := fmt.Sprintf("insert into testing(name) values('%v')",name)
	_, err := a.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

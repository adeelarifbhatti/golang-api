package main

import(
	"log"
	"testing"
	"fmt"
	"net/http/httptest"
	"net/http"
)
var a App

func TestMain(m *testing.M) {
	err := a.Start("root",DbPassword,DBName)
	if err != nil {
		log.Fatal(err)
	}
	createDB()
	m.Run()
}
func createDB() {
	createDatabase := `CREATE DATABASE testing;`
	useDatabase := `use testing;`
	createTable := `CREATE TABLE languages(id int Primary Key NOT NULL AUTO_INCREMENT,name VARCHAR(255));`
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
func clearTable() {
	a.DB.Exec("delete from testing")
	a.DB.Exec("Alter table testing Auto_Increment=1")
}
func addLanguage(name string){
	clearTable()
	query := fmt.Sprintf("insert into languages(name) values('%v')", name)
	_,err := a.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
func TestGetLanguage(t *testing.T) {
	clearTable()
	addLanguage("Java")
	req, _ := http.NewRequest("GET","/language/1",nil)
	response := sendRequest(req)
	checkStatusCode(t, http.StatusOK,response.Code)
}

func checkStatusCode(t *testing.T, expect int, gotten int){
	if expect != gotten {
		t.Errorf("Expected %v, Got %v", expect,gotten)
	}
	if expect == gotten {
	fmt.Println("expected was ### ", expect, " and what we got is ### " ,gotten)
	}
}
func sendRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)
	return recorder
}
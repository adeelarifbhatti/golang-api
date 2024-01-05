package main

import(
	"log"
	"testing"
	"fmt"
	"net/http/httptest"
	"net/http"
	"bytes"
	"encoding/json"
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
func TestCreateLanguage(t *testing.T){
	clearTable()
	var language = []byte(`{"name": "Python"}`)
	req, _ := http.NewRequest("POST","/language", bytes.NewBuffer(language))
	req.Header.Set("Content-Type","application/json")
	response := sendRequest(req)
	checkStatusCode(t, http.StatusCreated,response.Code)
	var m map[string]interface{}
	var lang map[string]interface{}
	json.Unmarshal(response.Body.Bytes(),&m)
	json.Unmarshal(language, &lang)
	if m["name"] != lang["name"] {
		t.Error("Expected name ","Python", " Gotten is ", m["name"])
	}

}

func checkStatusCode(t *testing.T, expect int, gotten int){
	if expect != gotten {
		t.Errorf("Expected %v, Got %v", expect,gotten)
	}
	if expect == gotten {
	fmt.Print(" Success ! \n expected was ### ", expect, " and what we got is ### " ,gotten)
	}
}
func sendRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)
	return recorder
}
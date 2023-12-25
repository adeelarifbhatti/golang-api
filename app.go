package main

import  (
		"database/sql" 
		"net/http"
		"github.com/gorilla/mux"
		_ "github.com/go-sql-driver/mysql"
		"log"
		"fmt"
		)
type App struct {
	Router *mux.Router
	DB *sql.DB
}
func checkError(e error){
	if e!= nil {
		log.Fatalln(e)
	}
}
func (app *App) Start() error {
	var err error
	connectString := fmt.Sprintf("%v:%v@tcp(mysql:3306)/%v",DbUser,DbPassword,DBName)
	db, err := sql.Open("mysql", connectString)
	fmt.Println(db)
	checkError(err)
	app.Router = mux.NewRouter().StrictSlash(true)
	return nil
}
func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address,app.Router))
}
func (app *App) handleRoute() {
	// app.Router.HandleFunc("/language", getLanguage).Methods("GET")
}

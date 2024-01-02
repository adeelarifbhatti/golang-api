package main

import  (
		"database/sql" 
		"net/http"
		"github.com/gorilla/mux"
		_ "github.com/go-sql-driver/mysql"
		"log"
		"fmt"
		"strconv"
		"encoding/json"
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
	app.DB, err = sql.Open("mysql", connectString)
	fmt.Println(app.DB)
	checkError(err)
	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoute()
	return nil
}
func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address,app.Router))
}
func (app *App) handleRoute() {
	 app.Router.HandleFunc("/languages", app.getLanguages).Methods("GET")
	 app.Router.HandleFunc("/language/{id}", app.getLanguage).Methods("GET")
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}){
	fmt.Println("Payload Unmarshal " , payload)
	response, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	fmt.Println("From sendResponse app ",response)
	w.Write(response)
}
func sendError(w http.ResponseWriter, statusCode int, err string){
	error_message := map[string]string{"error": err}
	sendResponse(w,statusCode,error_message)
}
func (app *App) getLanguages(w http.ResponseWriter, r *http.Request){
	languages, err := getLanguages(app.DB)
	if err != nil {
		sendError(w,http.StatusInternalServerError,err.Error())
		return
	}
	fmt.Println("From getLanguages from app  ", languages)
	sendResponse(w,http.StatusOK,languages)
}
func (app *App) getLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w,http.StatusBadRequest,err.Error())
		return
	}	
	lang := language{id: id}
	err = lang.getLanguage(app.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, "Language not found")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	fmt.Print("\n from app getLanguage  ",lang, "  and  ", lang.id, "  and ", id,"\n")
	fmt.Println(err)
	sendResponse(w,http.StatusOK,lang)
}
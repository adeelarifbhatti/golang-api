package main
import ("fmt"
		"net/http"
		"log"
		"encoding/json"
		"reflect"
		"strconv"
		)
type Language struct {
	name string
	id int
}
var languages []Language

func mainpage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome")
	log.Println("Docker changes")
}
func returnLanguages(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Println("returnLanguages")
	json.NewEncoder(w).Encode(languages)
}
func startServer(){
	http.HandleFunc("/",mainpage)
	http.HandleFunc("/languages",returnLanguages)
	http.ListenAndServe(":8080",nil)
}
func main(){
	languages = []Language{
		Language{name:"Golang", id: 1},
		Language{name:"Java", id: 2},
		Language{name:"Python", id: 3},
		Language{name:"Kotin", id: 4},
	}
	startServer()
}

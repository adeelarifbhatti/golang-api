package main
import ("fmt"
		"net/http"
		"log"
		"encoding/json"
		"reflect"
		"strconv"
		"github.com/gorilla/mux"
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
	log.Println("returnLanguages")
	json.NewEncoder(w).Encode(languages)
}
func returnLanguage(w http.ResponseWriter, r *http.Request){
	// Changing following because of gorilla/mux
	// log.Println("returnLanguage")
	// log.Println(r.URL.Path)
	// key := r.URL.Path[len("/language/"):]
	// log.Println("key is ",key)
	vars := mux.Vars(r)
	key := vars["id"]
	for _, languages := range languages {
		// log.Println("Type of key is: \n", reflect.TypeOf(key))
		id,err := strconv.Atoi(key)
		log.Println("err in strconv.Atoi(languages.id) is ", err," id is ", reflect.TypeOf(id))
		// log.Println("Type of key is: \n", reflect.TypeOf(key))
		if languages.id == id {
			fmt.Println("languages.id ",languages.id)
			fmt.Println("languages.name ",languages.name)
			fmt.Println(" languages.id ", languages.id)
			json.NewEncoder(w).Encode(languages.id)
			fmt.Println("languages.name ", languages.name)
			fmt.Println("language json response  ",languages)
			json.NewEncoder(w).Encode(languages.name)
			// w.Header().Set("Content-Type", "application/json")
			// w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(languages)
			fmt.Println("languages.id ",languages.id)
		}
	}
	// log.Println(languages)
	// log.Println(json.NewEncoder(w).Encode(languages))
	// json.NewEncoder(w).Encode(languages)
}
func startServer(){
	// replacing following with gorilla/mux
	// http.HandleFunc("/",mainpage)
	// http.HandleFunc("/languages",returnLanguages)
	// http.HandleFunc("/language/",returnLanguage)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",mainpage)
	myRouter.HandleFunc("/languages",returnLanguages)
	myRouter.HandleFunc("/language/{id}",returnLanguage)
	http.ListenAndServe(":8080",myRouter)
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

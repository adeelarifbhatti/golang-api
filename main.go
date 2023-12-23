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
	log.Println("returnLanguages")
	json.NewEncoder(w).Encode(languages)
}
func returnLanguage(w http.ResponseWriter, r *http.Request){
	log.Println("returnLanguage")
	log.Println(r.URL.Path)
	key := r.URL.Path[len("/language/"):]
	log.Println("key is ",key)
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
	http.HandleFunc("/",mainpage)
	http.HandleFunc("/languages",returnLanguages)
	http.HandleFunc("/language/",returnLanguage)
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

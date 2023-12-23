package main
import ("fmt"
		"net/http"
		)
func mainpage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome")
	fmt.Println("Docker changes")
}
func main(){
	http.HandleFunc("/",mainpage)
	http.ListenAndServe(":8080",nil)
}

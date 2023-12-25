package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
type language struct {
	id int
	name string
}
func checkError(e error){
	if e!= nil {
		log.Fatalln(e)
	}
}

func main(){
	connectString := fmt.Sprintf("%v:%v@tcp(mysql:3306)/%v",DbUser,DbPassword,DBName)
	db, err := sql.Open("mysql", connectString)
	checkError(err)
	// var id int
	// var name string
	// fmt.Print("Type a number")
	// fmt.Scan(&id)
	// fmt.Print("Type a name")
	// fmt.Scan(&name)
	result, err := db.Exec("insert into languages values(10,'kotlin')")
	checkError(err)
	lastInsertId, err := result.LastInsertId()
	fmt.Println("lastInsertId is ",lastInsertId)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected is ", rowsAffected)
	rows, err := db.Query("select * from languages")
	fmt.Println("err is   " ,err)
	var lang language
	for rows.Next(){
		err := rows.Scan(&lang.id,&lang.name)
		fmt.Println("err is   " ,err)
		fmt.Println(lang)
	}
	defer db.Close()
}
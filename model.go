package main

import (
		"database/sql"
		"fmt"
	)
type language struct {
	id int `json:"id"`
	name string `json:"name"`
}
func getLanguages(db *sql.DB) ([]language,error){
	query := "select * from languages"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	lang := []language{}
	for rows.Next(){
		var l language
		err := rows.Scan(&l.id, &l.name)
		fmt.Println("From get Languages in Model ",l)
		if err != nil {
			return nil, err	
		}
		lang = append(lang,l)
	}
	fmt.Println(lang)
	return lang, nil
}

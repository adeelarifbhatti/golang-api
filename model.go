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
func (l *language) getLanguage(db *sql.DB) error {
	fmt.Println("inside App getLanguage")
	query := fmt.Sprintf("select name from languages where id=%v", l.id)
	row := db.QueryRow(query)
	err := row.Scan(&l.name)
	fmt.Println("\nfrom model getLanguage  ", l.name ,"\n")
	if err != nil {
		return err
	}
	return err
}
func (lang *language) createLanguage(db *sql.DB) error {
	query := fmt.Sprintf("insert into languages(name) values('%v')",lang.name)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	lang.id = int(id)
	return nil
}

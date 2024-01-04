package main

import (
		"database/sql"
		"fmt"
		"errors"
	)
type language struct {
	ID int `json:"id"`
	NAME string `json:"name"`
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
		err := rows.Scan(&l.ID, &l.NAME)
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
	query := fmt.Sprintf("select name from languages where id=%v", l.ID)
	row := db.QueryRow(query)
	err := row.Scan(&l.NAME)
	fmt.Println("\nfrom model getLanguage  ", l.NAME )
	if err != nil {
		return err
	}
	return err
}
func (lang *language) createLanguage(db *sql.DB) error {
	query := fmt.Sprintf("insert into languages(name) values('%v')",lang.NAME)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	lang.ID = int(id)
	return nil
}
func (lang *language) updateLanguage(db *sql.DB) error {
	query := fmt.Sprintf("update languages set name='%v' where id=%v",lang.NAME, lang.ID)
	result,err := db.Exec(query)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("result.RowsAffected() >> ", rowsAffected)
	if rowsAffected == 0 {
		return errors.New("No such rows exists")
	}
	return err
}
func (lang *language) deleteLanguage(db *sql.DB) error {
	query := fmt.Sprintf("delete from languages where id=%v", lang.ID)
	_,err := db.Exec(query)
	return err
}

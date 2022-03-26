package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID    string
	title string
	body  string
}

func get_posts() {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	get, err := db.Query("select id,title,body from posts order by id")
	if err != nil {
		panic(err)
	}
	for get.Next() {
		var post Post
		err = get.Scan(&post.ID, &post.title, &post.body)
		if err != nil {
			panic(err)
		}
		fmt.Print(post.ID, ": ", post.title, ": ", post.body)
		fmt.Println()
	}
	defer db.Close()
}
func insert_post(post Post) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	insert, err := db.Query("insert into posts(title,body) Values (?,?)", post.title, post.body)
	if err != nil {
		panic(err)
	}
	insert.Close()
	defer db.Close()
}

func delete_post(id string) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	delete, err := db.Query("delete from posts where id= ?", id)
	if err != nil {
		panic(err)
	}
	delete.Close()
	defer db.Close()
}
func update_post(id string, post Post) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	update, err := db.Query("update posts set title=?,body=? where id= ?", post.title, post.body, id)
	if err != nil {
		panic(err)
	}
	update.Close()
	defer db.Close()
}
func main() {
}

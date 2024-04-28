package models

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

type Post struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func setupDb() *sql.DB {

	connstr := "postgres://tom:Motionless1nWhite18962@192.168.1.20:5432/blogdb?sslmode=disable"

	db, err := sql.Open("postgres", connstr)

	if err != nil {
		panic(err)
	}

	fmt.Print("db connection sucessful")
	return db

}

func GetPosts() []Post {
	db := setupDb()

	query := "SELECT * FROM posts"

	var posts []Post
	var post Post
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)

	}

	defer rows.Close()
	defer db.Close()
	return posts

}

func GetPost(id uint64) Post {

	db := setupDb()

	var p Post

	statement := "SELECT * FROM posts WHERE ID  = $1"

	rows, err := db.Query(statement, id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&p.ID, &p.Title, &p.Body)
		if err != nil {
			log.Fatal(err)
		}

	}

	defer rows.Close()
	defer db.Close()
	return p

}

func CreatePost(post Post) {
	db := setupDb()
	defer db.Close()
	statement := "INSERT INTO posts(TITLE,BODY) VALUES($1,$2)"
	db.Exec(statement, post.Title, post.Body)
}

func UpdatePost(post Post) {
	db := setupDb()
	defer db.Close()

	statment := "UPDATE posts SET TITLE = $1, BODY = $2 WHERE ID = $3"

	_, err := db.Exec(statment, post.Title, post.Body, post.ID)
	if err != nil {
		log.Fatal(err)
	}

}

func DeletePost(id uint64) {
	db := setupDb()
	defer db.Close()

	statement := "DELETE FROM posts WHERE ID = $1"
	res, err := db.Exec(statement, id)
	if err != nil {
		log.Fatal(err)
	}

	res.RowsAffected()

}

// TOOD: Make makefile

package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Tweets struct {
	Id    int
	Tweet string
}

var DbConnection *sql.DB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func getPostTweet(w http.ResponseWriter, r *http.Request) {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	v := r.FormValue("tweet")
	cmd := `INSERT INTO tweets(tweet)VALUES(?)`
	DbConnection.Exec(cmd, v)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tweet/", getPostTweet)
	http.ListenAndServe(":8080", nil)
}

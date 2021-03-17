package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}

type Articles []Article

var articles = Articles{
	Article{Title: "Pertama", Desc: "Desc nya"},
	Article{Title: "Kedua", Desc: "Desc dua"},
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/post-article", withLogging(postArticle))
	http.ListenAndServe(":3000", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kamu ada dihome"))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newArticle Article
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Can't read body", http.StatusInternalServerError)
		// }
		// w.Write([]byte(string(body)))

		err := json.NewDecoder(r.Body).Decode(&newArticle)

		if err != nil {
			fmt.Println("Ada error", err)
		}

		articles = append(articles, newArticle)

		json.NewEncoder(w).Encode(articles)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged koneksi dari", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}

}

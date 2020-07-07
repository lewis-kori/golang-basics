package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article is an instance of a blog
type Article struct {
	Id          string `json:"Id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

var articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	// capture the body of our request

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func updateEntireArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	for index, article := range articles {
		if article.Id == key {

			json.Unmarshal(reqBody, &article)
			articles[index] = article
			json.NewEncoder(w).Encode(article)
		}
	}

}

// func patchArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	for index, article := range articles {
// 		if article.Id == key {

// 		}
// 	}
// }

// func (a Article) keys() string {
// 	v := reflect.ValueOf(a)
// 	for i := 0; i < v.NumField(); i++ {
// 		return v.Type().Field(i).Name
// 	}
// }

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for index, article := range articles {
		if article.Id == key {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/articles", returnAllArticles)

	// keep note of the ordering
	myRouter.HandleFunc("/article/", createArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", updateEntireArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	articles = []Article{
		{Id: "1", Title: "Hello", Description: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Description: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

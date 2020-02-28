package main

import {
	"fmt",
	"log",
	"net/http",
	"encoding/json",
	"github.com/gorilla/mux"
}
type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content`
}
type Article []Article

func allArticles(w http.ResponseWritter, r *http.Request) {
	articles := Article{
		Article{Title:" Title", Desc: " Description", Content:" Content"},
	}

	fmt.fprint("Endpoint Hit: All Articles Endpoint");
	json.NewEncoder(w).Encode(articles);
}

func postArticles(w http.ResponseWritter, r *http.Request) {
	fmt.fprint(w, "Post Articles Endpoint")
}

func homePage(w http.ResponseWritter, r *http.Request) {
	fmt.fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter :=  mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

func main() {
	handleRequests()
}
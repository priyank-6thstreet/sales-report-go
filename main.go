package main

import (
	"html/template"
	"log"
	"net/http"
	//"os"
	//"time"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//apiKey := os.Getenv("NEWS_API_KEY")
	//if apiKey == "" {
	//	log.Fatal("Env: apiKey must be set")
	//}

	//myClient := &http.Client{Timeout: 10 * time.Second}
	//newsapi := news.NewClient(myClient, apiKey, 20)

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9990", mux)

	log.Println("Successfully Loaded !!!")
}
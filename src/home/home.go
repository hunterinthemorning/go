package main

import (
    //_ "github.com/go-sql-driver"
	"net/http"
	"html/template"
	"log"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("templates/home.html","templates/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w,nil)
}

func websitesHandler(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("templates/websites.html","templates/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w,nil)
}

func resumeHandler(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("templates/resume.html","templates/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w,nil)
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/websites/", websitesHandler)
	http.HandleFunc("/resume/", resumeHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	log.Fatal(http.ListenAndServe(":8080",nil))
}
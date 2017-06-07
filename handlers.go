package main

import (
	"net/http"
	"fmt"
	"html"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Dict API is Running is running on %q</h1>", html.EscapeString(port));
}

func GetAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, RawWords)
}

func GetOne(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, SingleWordAsString())
}

func Define(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	rWord := vars["word"]
	fmt.Fprintf(w, DefineSingleWord(rWord))
}
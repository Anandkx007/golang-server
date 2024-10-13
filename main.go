package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}

	fmt.Fprintf(w, "Post resquest success\n")

	name := r.FormValue("name")

	fmt.Fprintf(w, "Name = %s\n", name)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {

		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {

		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello! Anand")

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8084 \n")

	if err := http.ListenAndServe(":8084", nil); err != nil {
		log.Fatal(err)
	}
}

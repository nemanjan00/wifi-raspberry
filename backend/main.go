package main;

import (
	"fmt"
	"log"
	"net/http"
);

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]);
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/api", handler);
	log.Fatal(http.ListenAndServe(":8080", nil));
}



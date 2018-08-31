package main;

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"bytes"

	"./config"
);

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]);
}

func getListenerParam() (string){
	port := strconv.Itoa(config.GetPort());
	hostname := config.GetHostname();

	var param bytes.Buffer;
	param.WriteString(hostname);
	param.WriteString(":");
	param.WriteString(port);

	return param.String();
}

func main() {
	fs := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/", fs);

	http.HandleFunc("/api/", handler);

	param := getListenerParam();

	fmt.Printf("Starting wifi-raspberry on port %s.\n", param);
	log.Fatal(http.ListenAndServe(param, nil));
}


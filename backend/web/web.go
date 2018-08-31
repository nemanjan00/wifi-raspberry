package web;

import (
	"../config"

	"fmt"
	"log"
	"strconv"
	"bytes"

	"os"
	"net/http"
);

func getListenerParam() (string){
	port := strconv.Itoa(config.GetPort());
	hostname := config.GetHostname();

	var param bytes.Buffer;
	param.WriteString(hostname);
	param.WriteString(":");
	param.WriteString(port);

	return param.String();
}

func Start(){
	param := getListenerParam();

	fmt.Printf("Starting wifi-raspberry on port %s.\n", param);
	log.Fatal(http.ListenAndServe(param, nil));
}

func ConfigureServer(){
	if _, err := os.Stat("./build/static"); os.IsNotExist(err) {
		fs := http.FileServer(http.Dir("./static"))
		http.Handle("/", fs);
	} else {
		fs := http.FileServer(http.Dir("./build/static"))
		http.Handle("/", fs);
	}
}


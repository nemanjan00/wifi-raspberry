package config;
	
import (
	"os"
	"strconv"
);

func GetPort() (int){
	port := os.Getenv("PORT");

	portInt, err := strconv.Atoi(port);

	if(err == nil){
		return portInt;
	} else {
		return 3000;
	}
}

func GetHostname() (string){
	host := os.Getenv("HOST");

	if(len(host) > 0){
		return host;
	} else {
		return "127.0.0.1";
	}
}


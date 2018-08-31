package main;

import (
	"./web"
);

func main() {
	web.ConfigureServer();
	web.Start();
}


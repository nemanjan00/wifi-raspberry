wifi-raspberry: 
	go build -o ./build/wifi-raspberry ./backend/main.go

watch: 
	nodemon -e go -w ./backend -x "go run" backend/main.go

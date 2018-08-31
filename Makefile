backend: 
	go build -o ./build/wifi-raspberry ./backend/main.go

watch-backend: 
	nodemon -e go -w ./backend -x "go run" backend/main.go

backend: 
	go build -o ./build/wifi-raspberry ./backend/main.go

watch-backend: 
	nodemon --config ./backend.nodemon.json

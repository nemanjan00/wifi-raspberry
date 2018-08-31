all: build/wifi-raspberry frontend/static/main.js

build/wifi-raspberry:
	go build -o ./build/wifi-raspberry ./backend/main.go

watch-backend: 
	nodemon --config ./backend.nodemon.json

frontend/static/main.js:
	webpack-cli --config ./frontend/webpack.config.js

watch-frontend:
	webpack-cli --watch --config ./frontend/webpack.config.js


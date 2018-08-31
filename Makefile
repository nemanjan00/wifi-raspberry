all: build/wifi-raspberry build/static/main.js

build/wifi-raspberry:
	go build -o ./build/wifi-raspberry ./backend/main.go

watch-backend: 
	nodemon --config ./devfiles/backend.nodemon.json

build/static/main.js:
	webpack-cli --config ./frontend/webpack.config.js

watch-frontend:
	webpack-cli --watch --config ./frontend/webpack.config.js


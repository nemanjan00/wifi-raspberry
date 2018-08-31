backend:
	go build -o ./build/wifi-raspberry ./backend/main.go

watch-backend: 
	nodemon --config ./backend.nodemon.json

frontend:
	webpack-cli --config ./frontend/webpack.config.js

watch-frontend:
	webpack-cli --watch --config ./frontend/webpack.config.js


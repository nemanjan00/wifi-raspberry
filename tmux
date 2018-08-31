rename-session cv
send "make watch-backend" C-m
new-window
send "webpack-cli --watch --config ./frontend/webpack.config.js" C-m
new-window
send "vim ." C-m


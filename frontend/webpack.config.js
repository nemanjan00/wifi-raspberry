'use strict'

const { VueLoaderPlugin } = require('vue-loader')

const path = require("path");
const _root = path.resolve(__dirname, '..');

function root(args) {
	args = Array.prototype.slice.call(arguments, 0);
	return path.join.apply(path, [_root].concat(args));
}

module.exports = {
	output: {
		path: root('frontend/static'),
		publicPath: '/',
		filename: '[name].js'
	},
	mode: 'development',
	entry: [
		'./frontend/app.js'
	],
	module: {
		rules: [
			{
				test: /\.vue$/,
				use: 'vue-loader'
			},
			{
				test: /\.css$/,
				use: [ 'style-loader', 'css-loader' ]
			}
		]
	},
	plugins: [
		new VueLoaderPlugin()
	]
}


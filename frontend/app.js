const Vue = require("vue").default;
const App = require("./App.vue").default;

new Vue({
	el: '#app',
	render: h => h(App)
});


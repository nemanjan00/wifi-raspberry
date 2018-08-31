//import 'bootswatch/dist/superhero/bootstrap.css';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

import './app.css';

import '@fortawesome/fontawesome-free/js/all.js';

import Vue from 'vue';
import App from './App.vue';

import BootstrapVue from 'bootstrap-vue'

Vue.use(BootstrapVue);

new Vue({
	el: '#app',
	render: h => h(App)
});


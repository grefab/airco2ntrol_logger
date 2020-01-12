import Vue from 'vue'

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue);
Vue.config.productionTip = false;

import VueC3 from 'vue-c3'

import App from './App.vue'

new Vue({
  render: h => h(App),
}).$mount('#app');

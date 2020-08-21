import Vue from 'vue'
import Element from 'element-ui'
import App from './App.vue'
import router from './router'
import store from './store'
import api from './services/api'
import 'element-ui/lib/theme-chalk/index.css';
import './index.css';

Vue.config.productionTip = false
Vue.use(Element)
Vue.use(api);

Vue.config.errorHandler = function (err, vm, info) {
  console.error(err);
  Vue.prototype.$message.error(err.toString())
};

new Vue({
  router,
  store,
  render: function (h) { return h(App) }
}).$mount('#app')

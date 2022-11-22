import Vue from 'vue';
import app from "./App.vue"
import VueRouter from 'vue-router';
import axios from "axios";
import VueAxios from "vue-axios";
import VueCookies from 'vue-cookies-reactive';
import vSelect from "vue-select";
export default router

import HomeView from './views/HomeView.vue'
import LoginView from './views/LoginView.vue'
import RegisterView from './views/RegisterView.vue'
import NewThreadView from './views/NewThreadView.vue'
import GlobalStores from './plugins/GlobalStores';


Vue.config.productionTip = false;

const store = Vue.observable({ 
  auth: {
    user: '',
    set(unam){
      store.auth.user = unam;
    }
  }
});

Vue.use(GlobalStores)
Vue.component("v-select", vSelect);

const client = axios.create({
  baseURL: "/api/v1",
});

Vue.use(VueCookies);
Vue.use(VueAxios, client);



const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [{
        path: '/',
        alias: ['/index.html'],
        name: '',
        component: HomeView
      },
      {
        path: '/login',
        name: 'login',
        component: LoginView
      },
      {
        path: '/register',
        name: 'register',
        component: RegisterView
      },
      {
        path: '/newThread',
        name: 'newThread',
        component: NewThreadView
      },
      {
        path: '*',
        name: '404',
        component: () => import('./views/NotFound.vue'),
      },],
});

Vue.use(VueRouter);

new Vue({
  router: router,
  render: h => h(app)
}).$mount('#app');
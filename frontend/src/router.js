import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from './components/Home.vue'
import Login from "./components/login/Login";
import EditAudio from "./components/audios/EditAudio";
import Audios from "./components/audios/Audios";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        component: Home,
    },
    {
        path: '/edit',
        component: EditAudio
    },
    {
      path: '/login',
      component: Login
    }
];

const router = new VueRouter({
    routes
});

export default router

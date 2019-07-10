import Vue                from 'vue';
import VueRouter          from 'vue-router';
import { createProvider } from './vue-apollo';
import ListMovie          from './components/ListMovie';
import DetailMovie        from './components/DetailMovie';
import DetailUser         from './components/DetailUser';
import App                from './App.vue';

Vue.use(VueRouter);

Vue.config.productionTip = false;

new Vue({
    apolloProvider: createProvider(),
    render        : h => h(App),
    router        : new VueRouter({
        mode  : 'history',
        routes: [
            { path: '/', component: ListMovie },
            { path: '/movie/:id', component: DetailMovie },
            { path: '/user/:id', component: DetailUser },
        ],
    }),
}).$mount('#app');

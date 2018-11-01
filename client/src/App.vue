<template>
    <div id="app">
        <img alt="Vue logo" src="./assets/logo.png">
        {{JSON.stringify(users)}}
    </div>
</template>

<script>
    import gql               from 'graphql-tag';
    import HelloWorld        from './components/HelloWorld.vue';
    import { axiosProvider } from './services';

    export default {
        name        : 'app',
        components  : {
            HelloWorld,
        },
        data() {
            return { users: [],loading: 0, };
        },
        beforeCreate: function() {
            axiosProvider.request.get('/api').then(response => this.helloText = response.data);
        },
        apollo      : {
            users: {
                query: gql`{users {name}}`,
                loadingKey: 'loading',
            },
        },
    };
</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
    }
</style>

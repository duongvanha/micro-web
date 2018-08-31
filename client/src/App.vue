<template>
    <div id="app">
        <ListMovie/>
        <MovieDetail/>
    </div>
</template>

<script>
    import HelloWorld        from './components/HelloWorld.vue';
    import { axiosProvider } from './services';
    import ListMovie         from './components/ListMovie';
    import MovieDetail       from './components/MovieDetail';

    export default {
        name        : 'app',
        components  : {
            MovieDetail,
            ListMovie,
            HelloWorld,
        },
        data() {
            return {helloText: ''}
        },
        beforeCreate: function () {
            axiosProvider.request.get('/api').then(response => this.helloText = response.data)
        },
    }
</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>

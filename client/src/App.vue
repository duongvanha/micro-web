<template>
    <div id="app">
        <ListMovie v-on:clickMovie="selectMovie"/>
        <MovieDetail v-bind:movie="movie"/>
    </div>
</template>

<script>
    import ListMovie   from './components/ListMovie';
    import MovieDetail from './components/MovieDetail';

    export default {
        name      : 'app',
        components: {
            MovieDetail,
            ListMovie,
        },
        data() {
            return { movie: null };
        },
        methods   : {
            selectMovie(movie) {
                this.movie = movie;
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
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>

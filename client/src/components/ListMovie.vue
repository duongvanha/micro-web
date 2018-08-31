<template>
    <ul>
        <li v-for="(item) in movies" v-on:click="$emit('clickMovie',item)">
            {{item.namevi}}
        </li>
    </ul>
</template>

<script>

    import { axiosProvider } from '../services';

    export default {
        name        : 'ListMovie',
        data() {
            return { movies: [] };
        },
        beforeCreate: function() {
            axiosProvider.request.get('/api/movies').then(response => this.movies = response.data);
        },
    };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    ul {
        list-style: none;
        padding: 20px;
        height: 100%;
        overflow: scroll;
    }
</style>

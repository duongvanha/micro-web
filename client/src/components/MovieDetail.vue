<template>
    <vue-plyr>
        <video v-bind:poster="movie.image" src="video.mp4" v-if="movie">
            <source v-if="detailMovie" v-for="item in detailMovie.medias" v-bind:src="item.url" type="video/mp4"
                    v-bind:size="item.resolution"/>
        </video>
    </vue-plyr>
</template>

<script>

    import { axiosProvider } from '../services';

    export default {
        name : 'MovieDetail',
        props: ['movie'],
        data() {
            return { detailMovie: { medias: [] } };
        },
        watch: {
            movie(val, oldval) {
                axiosProvider.request.post('/api/movie', { url: this.movie.url })
                    .then(response => {
                        this.detailMovie = response.data;
                    });
            },
        },
    };
</script>

<style scoped>

</style>

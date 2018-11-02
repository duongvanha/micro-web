<template>
    <div>
        Name en {{JSON.stringify(movie)}}
        Name vi
    </div>
</template>

<script>
    import movieQuery from '../graphql/movie.graphql';

    export default {
        name  : 'DetailMovie',
        data() {
            return { movieQuery, movie: {} };
        },
        apollo: {
            movie: {
                query: movieQuery,
                variables() {
                    return { url: `http://www.phimmoi.net/phim/${this.$route.params.id}/` };
                },
            },
        },
    };
</script>

<style scoped>

</style>

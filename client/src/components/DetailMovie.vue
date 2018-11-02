<template>
    <div>
        Name en {{movie.name_en}}
        Name vi {{movie.name_vi}}
        View {{movie.view}}
    </div>
</template>

<script>
    import movieQuery         from '../graphql/movie.graphql';
    import numberViewMutation from '../graphql/updateNumberView.graphql';

    export default {
        name    : 'DetailMovie',
        data() {
            return { movieQuery, movie: {} };
        },
        methods : {
            updateNumberView() {
                return this.$apollo.mutate({
                    mutation : numberViewMutation,
                    variables: {
                        url: this.url,
                    },
                });
            },
        },
        computed: {
            url() {
                return `http://www.phimmoi.net/phim/${this.$route.params.id}/`;
            },
        },
        beforeMount() {
            this.updateNumberView();
        },
        apollo  : {
            movie: {
                query: movieQuery,
                variables() {
                    return { url: this.url };
                },
            },
        },
    };
</script>

<style scoped>

</style>

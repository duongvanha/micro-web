<template>
    <div>
        <ul>
            <li v-for="movie in movies.collection">
                <router-link :to="movie.url | urlMovie">
                    {{movie.name_en}} - {{movie.name_vi}} - {{movie.view}}
                </router-link>
            </li>
        </ul>
        <div class="actions">
            <button v-if="hasNext" @click="loadMore">Show more</button>
        </div>
    </div>
</template>

<script>
    import movies from '../graphql/movies.graphql';

    export default {
        name   : 'ListMovie',
        data() {
            return { movies: [], page: 0, hasNext: true, pageSize: 10 };
        },
        filters: {
            urlMovie(url) {
                const resultsName = new RegExp('http://www\.phimmoi\.net/phim/(.*)/').exec(url);
                if (!resultsName) {
                    return '#';
                }
                return `/movie/${resultsName[1]}`;
            },
        },
        methods: {
            loadMore() {
                // Fetch more data and transform the original result
                this.$apollo.queries.movies.fetchMore({
                    // New variables
                    variables  : {
                        page    : this.page + 1,
                        pageSize: 10,
                    },
                    // Transform the previous result with new data
                    updateQuery: (previousResult, { fetchMoreResult }) => {
                        const collection = fetchMoreResult.movies.collection;
                        const hasNext    = fetchMoreResult.movies.hasNext;
                        const page       = fetchMoreResult.movies.page;

                        this.hasNext = hasNext;
                        this.page    = page;

                        return {
                            movies: {
                                __typename: previousResult.movies.__typename,
                                collection: [...previousResult.movies.collection, ...collection],
                                page,
                                hasNext,
                            },
                        };
                    },
                });
            },
        },
        apollo : {
            movies: {
                query    : movies,
                variables: {
                    page    : 0,
                    pageSize: 10,
                },
            },
        },
    };
</script>

<style scoped>

</style>

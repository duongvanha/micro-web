import { GraphQLServer } from 'graphql-yoga';
import Knex              from 'knex';
import typeDefs          from './typeDefs.graphql';
import uuid              from 'uuid';
import ip                from 'ip';
import Consul            from 'consul';

const IP   = ip.address();
const PORT = 4000;

const CONSUL_ID = uuid.v4();

const consul = Consul({ port: 8501, IP: '127.0.0.1' });

let details = {
    name   : 'data',
    address: IP,
    port   : PORT,
    id     : CONSUL_ID,
    check  : {
        ttl                              : '10s',
        deregister_critical_service_after: '1m',
    },
};

console.log(IP);

const knex = Knex({
    client    : 'pg',
    connection: {
        host    : process.env.POSTGRES_HOST,
        user    : process.env.POSTGRES_USER,
        password: process.env.POSTGRES_PASSWORD,
        database: process.env.POSTGRES_DB,
    },
});

const resolvers = {
    Query   : {
        categories: () => knex.select().table('categories').limit(10),
        users     : () => knex.select().table('users').limit(10),
        countries : () => knex.select().table('countries').limit(10),
        movies    : (parent, condition = { page: 1, pageSize: 10 }) => knex.select()
            .table('movies')
            .limit(condition.pageSize + 1)
            .offset(condition.page * condition.pageSize).then(data => {
                return ({
                    page      : condition.page,
                    hasNext   : data.length > condition.pageSize,
                    collection: data.splice(0, condition.pageSize),
                });
            }),
        movie     : (parent, condition = { url: '' }) => knex.select()
            .table('movies')
            .where('url', condition.url).then(collection => collection[0]),
    },
    Movie   : {
        directors(parent) {
            return knex.select()
                .table('users')
                .leftJoin('movie_director_user', 'users.href', 'movie_director_user.user_href')
                .where('movie_url', parent.url);
        },
        actors(parent) {
            return knex.select()
                .table('users')
                .leftJoin('movie_actor_user', 'users.href', 'movie_actor_user.user_href')
                .where('movie_url', parent.url);
        },
        categories(parent) {
            return knex.select()
                .table('categories')
                .leftJoin('user_categories', 'categories.href', 'user_categories.category_href')
                .where('movie_url', parent.url);
        },
        countries(parent) {
            return knex.select()
                .table('countries')
                .leftJoin('movie_country', 'countries.code', 'movie_country.country_code')
                .where('movie_url', parent.url);
        },
    },
    Category: {
        movies(parent) {
            return knex.select()
                .table('movies')
                .leftJoin('user_categories', 'movies.url', 'user_categories.movie_url')
                .where('category_href', parent.href).limit(10);
        },
    },
    User    : {
        actorMovies(parent) {
            return knex.select()
                .table('movies')
                .leftJoin('movie_actor_user', 'movie.url', 'movie_actor_user.movie_url')
                .where('user_href', parent.href);
        },
        directorMovies(parent) {
            return knex.select()
                .table('movies')
                .leftJoin('movie_director_user', 'movie.url', 'movie_director_user.movie_url')
                .where('user_href', parent.href);
        },
    },
    Mutation: {
        updateNumberView: async (data, condition) => {
            await knex.select()
                .table('movies')
                .where('url', condition.url)
                .increment('view', 1);

            return knex.select()
                .table('movies')
                .where('url', condition.url)
                .then(collection => collection[0]);

        },
    },
    // Country : {
    //     movies(parent) {
    //         return knex.select()
    //             .table('movie')
    //             .leftJoin('movie_country', 'movie.url', 'movie_country.movie_url')
    //             .where('country_code', parent.code).limit(10);
    //     },
    // },
};
const server    = new GraphQLServer({ typeDefs, resolvers });

(async function() {
    consul.agent.service.register(details, console.log);
    await server.start();
    console.log(`Server is running ip ${IP}:${PORT}`);
})();

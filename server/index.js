const PhimMoiProvider = require('./Movie/PhimMoi');
const express         = require('express');
const bodyParser      = require('body-parser');

const app           = express();
const movieProvider = new PhimMoiProvider();

const knex = require('knex')({
    client    : 'pg',
    connection: process.env.POSTGRES_URL + '?ssl=true',
});

app.use(bodyParser.json());

app.get('/', function(req, res) {
    res.send('hello word');
});

app.get('/movies', async (req, res) => {
    const page = req.query.page || 1;
    let total  = req.query.total || 10;
    let query  = knex.select('*').from('movies');
    query      = query.limit(total).offset(page * total);
    if (req.query.search) {
        query = query.where('nameEn', 'like', `%${req.query.search}%`).orWhere('nameVi', 'like', `%${req.query.search}%`);
    }

    res.json(await query);
});

app.post('/movie', (req, res) => {
    movieProvider.getMovieByUrl(`${req.body.url}/xem-phim.html`).then(movie => res.json(movie));
});

app.listen(3000, () => console.log('app listen port 3000'));

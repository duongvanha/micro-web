const express    = require('express');
const bodyParser = require('body-parser');
const cors       = require('cors');
const redis      = require('redis');
const {Pool}     = require('pg');


const app = express();

const pgClient = new Pool({
    user    : process.env.POSTGRES_USER,
    host    : process.env.POSTGRES_HOST,
    database: process.env.POSTGRES_DB,
    password: process.env.POSTGRES_PASSWORD,
});

pgClient.on('error', () => console.log('Lost PG connection'));

pgClient
    .query('CREATE TABLE IF NOT EXISTS values (number INT)')
    .catch(err => console.log(err));

const redisClient    = redis.createClient({
    host          : process.env.REDIS_HOST,
    retry_strategy: () => 1000,
});
const redisPublisher = redisClient.duplicate();

app.use(cors());
app.use(bodyParser.json());

app.get('/', function (req, res) {
    res.send('hello word');
});

app.get('/values/all', async (req, res) => {
    const values = await pgClient.query('SELECT * from values');

    res.send(values.rows);
});

app.get('/values/current', async (req, res) => {
    redisClient.hgetall('values', (err, values) => {
        res.send(values);
    });
});

app.post('/values', async (req, res) => {
    const index = req.body.index;

    if (parseInt(index) > 40) {
        return res.status(422).send('Index too high');
    }

    redisClient.hset('values', index, 'Nothing yet!');
    redisPublisher.publish('insert', index);
    pgClient.query('INSERT INTO values(number) VALUES($1)', [index]);

    res.send({working: true});
});

app.listen(3000, () => console.log('app listen port 3000'));
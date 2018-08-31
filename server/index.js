const express    = require('express');
const bodyParser = require('body-parser');

const app = express();

app.use(bodyParser.json());

app.get('/', function(req, res) {
  res.send('hello word');
});

app.get('/movies', async (req, res) => {
  const values = await pgClient.query('SELECT * from values');

  res.send(values.rows);
});

app.get('/movie/:id', async (req, res) => {
  redisClient.hgetall('values', (err, values) => {
    res.send(values);
  });
});

app.listen(3000, () => console.log('app listen port 3000'));

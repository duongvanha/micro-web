const Express = require('express');

const app = new Express();

app.get('/', function (req, res) {
    res.send('hello word');
});

app.listen(3000, () => console.log("app listen port 3000"));
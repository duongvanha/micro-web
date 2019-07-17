const express = require('express');
const uuid    = require('uuid');
const ip      = require('ip');
const Consul  = require('consul');


const app       = express();
const IP        = ip.address();
const PORT      = 4000;
const CONSUL_ID = uuid.v4();

const consul = Consul({ port: 8501, IP: '127.0.0.1', promisify: true });

const url = `http://${IP}:${PORT}/check`;

console.log(CONSUL_ID);
console.log(url);
const options = {
    name   : 'data',
    address: IP,
    port   : PORT,
    id     : CONSUL_ID,
    check  : {
        // note                             : 'HTTP API on port 4000',
        // 'http'  s                         : 'https://localhost:5000/health',
        // 'tls_skip_verify'                : false,
        // 'method'                         : 'POST',
        // 'header'                         : { 'x-foo': ['bar', 'baz'] },
        // 'interval'                       : '10s',
        // 'timeout'                        : '1s',
        // http                             : url,
        // ttl                              : '10s',
        deregister_critical_service_after: '1m',

        // name         : 'example',
        ttl: '10s',
        // http         : `${CONSUL_ID}@${IP}:${PORT}/`,
        // tlsskipverify: false,
        // serviceid    : CONSUL_ID,
        // id       : CONSUL_ID + 1,
        // notes        : 'This is an example check.',
    },
};

app.use(function(req, res, next) {
    console.log('LOGGED');
    next();
});

app.get('/check', (req, res) => {
    console.log('tick');
    res.send(`I'm ok`);
});

(async function() {
    await consul.agent.service.register(options);
    // await consul.agent.check.register({
    //     name     : 'example',
    //     ttl      : '15s',
    //     http     : url,
    //     serviceid: CONSUL_ID,
    //     id       : CONSUL_ID + 1,
    //     notes    : 'This is an example check.',
    // });

    // setInterval(() => {
    //     consul.agent.check.pass({ id: `service:${CONSUL_ID}` }).then(console.log).catch(console.error);
    //
    // }, 1000);

    const server = app.listen(PORT, () => console.log(`Example app listening on port ${PORT}!`));

    process.on('SIGINT', function() {
        server.close(() => console.log(' Stopping ...'));
        consul.agent.service.deregister(options);
        // consul.agent.check.deregister({
        //     id: CONSUL_ID + 1,
        // }, function(err) {
        //     if (err) throw err;
        // });
    });

})();

const express = require('express');
const uuid    = require('uuid');
const ip      = require('ip');
const Consul  = require('consul');
const os      = require('os');


console.log();


const app       = express();
const IP        = ip.address();
const PORT      = 4000;
const CONSUL_ID = uuid.v4();
const hostName  = os.hostname();

const consul = Consul({ port: 8500, host: 'consul', promisify: true });

const url = `http://${hostName}:${PORT}/check`;

console.log(CONSUL_ID);
console.log(url);

const options = {
    name   : 'data',
    address: IP,
    port   : PORT,
    id     : CONSUL_ID,
    check  : {
        // 'id'                             : 'api',
        'name'                           : 'HTTP API on port 5000',
        // 'http'                           : url,
        // 'tls_skip_verify'                : false,
        // 'method'                         : 'POST',
        // 'header'                         : { 'x-foo': ['bar', 'baz'] },
        // 'interval'                       : '10s',
        // 'timeout'                        : '1s'
        // note                             : 'HTTP API on port 4000',
        // 'http'  s                         : 'https://localhost:5000/health',
        // 'tls_skip_verify'                : false,
        // 'method'                         : 'POST',
        // 'header'                         : { 'x-foo': ['bar', 'baz'] },
        // 'interval'                       : '10s',
        // 'timeout'                        : '1s',
        // http                             : url,
        // ttl                              : '10s',
        deregister_critical_service_after: '1s',

        // name         : 'example',
        interval: '10s',
        http    : url,
        // tlsskipverify: false,
        // serviceid: CONSUL_ID,
        // id      : `${CONSUL_ID}:1`,
        notes   : 'Check service alive',
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
    const server = app.listen(PORT, () => console.log(`Example app listening on port ${PORT}!`));

    process.on('SIGINT', function() {
        consul.agent.service.deregister(options);
        server.close(() => console.log(' Stopping ...'));
    });

})();

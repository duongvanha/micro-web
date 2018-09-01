const redis = require('redis');

const redisClient = redis.createClient({
    host          : process.env.REDIS_HOST,
    retry_strategy: () => 1000,
});
const sub         = redisClient.duplicate();

function fib(index) {
    if (index < 2) return 1;
    return fib(index - 1) + fib(index - 2);
}

sub.on('message', (channel, message) => {
    console.log('go mesage', message);
    redisClient.hset('values', message, fib(parseInt(message)));
});

sub.subscribe('insert');
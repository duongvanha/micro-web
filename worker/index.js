const grpc  = require('grpc');
const proto = grpc.load('../proto/transfers.proto');

const server = new grpc.Server();

server.addService(proto.context.TransferService.service, {

    send(call, callback) {
        console.log(call.request);
        callback(null, { id: 1, text: 'demo' });
    },


});

server.bind('0.0.0.0:50050', grpc.ServerCredentials.createInsecure());

//Start the server
server.start();
console.log('grpc server running on port:', '0.0.0.0:50050');
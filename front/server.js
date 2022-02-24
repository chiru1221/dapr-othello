const express = require('express');
const path = require('path');
const request = require('request');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const BOARD_PROTO_PATH = '/root/src/proto/board.proto';
const CP_PROTO_PATH = '/root/src/proto/cp.proto';

const createPakageDefinition = (filename) => {
    return protoLoader.loadSync(
        filename,
        {keepCase: true,
         longs: String,
         enums: String,
         defaults: true,
         oneofs: true
    });
}

const board_proto = grpc.loadPackageDefinition(
    createPakageDefinition(BOARD_PROTO_PATH)
).board;
const cp_proto = grpc.loadPackageDefinition(
    createPakageDefinition(CP_PROTO_PATH)
).cp;

const boardClient = new board_proto.BoardApi('board:8080', grpc.credentials.createInsecure());
const cpClient = new cp_proto.CpApi('cp:5000', grpc.credentials.createInsecure());

const app = express();
app.use(express.json());
const port = 3000;

app.post('/putable', async (req, res) => {
    boardClient.putable(req.body, function(err, response) {
        res.send(response);
    });
})

app.post('/reverse', async (req, res) => {
    boardClient.reverse(req.body, function(err, response) {
        res.send(response);
    });
})

app.post('/cp', async (req, res) => {
    cpClient.attack(req.body, function(err, response) {
        res.send(response);
    })
})

// Serve static files
app.use(express.static(path.join(__dirname, 'othello/build')));

app.get('*', function (_req, res) {
    res.sendFile(path.join(__dirname, 'othello/build', 'index.html'));
});

app.listen(port, () => console.log(`Listening on port ${port}!`));

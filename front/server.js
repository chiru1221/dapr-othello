const express = require('express');
const path = require('path');
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

const app = express();
app.use(express.json());
const port = 3000;

const prod = process.env.NODE_ENV === 'production'
const daprPort = process.env.DAPR_GRPC_PORT || 50001;
const boardApi = (prod)? `localhost:${daprPort}`: 'board:8080';
const cpApi = (prod)? `localhost:${daprPort}`: 'cp:5000';
const boardClient = new board_proto.BoardApi(boardApi, grpc.credentials.createInsecure());
const cpClient = new cp_proto.CpApi(cpApi, grpc.credentials.createInsecure());
const boardMetadata = new grpc.Metadata();
boardMetadata.add('dapr-app-id', 'boardapi');
const cpMetadata = new grpc.Metadata();
cpMetadata.add('dapr-app-id', 'cpapi');

app.post('/putable', async (req, res) => {
    boardClient.putable({
        stone: req.body.stone,
        x: req.body.x,
        y: req.body.y,
        squares: req.body.squares,
    }, boardMetadata, function(err, response) {
        res.send(response);
    });
})

app.post('/reverse', async (req, res) => {
    boardClient.reverse({
        stone: req.body.stone,
        x: req.body.x,
        y: req.body.y,
        squares: req.body.squares,
    }, boardMetadata, function(err, response) {
        res.send(response);
    });
})

app.post('/cp', async (req, res) => {
    cpClient.attack({
        level: req.body.level,
        stone: req.body.stone,
        squares: req.body.squares,
    }, cpMetadata, function(err, response) {
        res.send(response);
    })
})

// Serve static files
app.use(express.static(path.join(__dirname, 'othello/build')));

app.get('*', function (_req, res) {
    res.sendFile(path.join(__dirname, 'othello/build', 'index.html'));
});

app.listen(port, () => console.log(`Listening on port ${port}!`));

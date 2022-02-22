const express = require('express');
const path = require('path');
const request = require('request');

const app = express();

const port = 3000;
const prod = process.env.NODE_ENV === 'production'
const daprPort = process.env.DAPR_HTTP_PORT || 3500;
const boardApi = (prod)? `http://localhost:${daprPort}/v1.0/invoke/boardapi/method`: 'http://board:8080';
const cpApi = (prod)? `http://localhost:${daprPort}/v1.0/invoke/cpapi/method`: 'http://cp:5000';

// The name of the state store is specified in the components yaml file. 
// For this sample, state store name is specified in the file at: https://github.com/dapr/quickstarts/blob/master/hello-kubernetes/deploy/redis.yaml#L4
// const stateStoreName = `statestore`;
// const stateUrl = `http://localhost:${daprPort}/v1.0/state/${stateStoreName}`;


app.post('/putable', async (req, res) => {
    // req.pipe(request(`http://board:8080/putable`)).pipe(res);
    // req.pipe(request(`${daprUrl}/boardapi/method/putable`)).pipe(res);
    req.pipe(request(`${boardApi}/putable`)).pipe(res);
})

app.post('/reverse', async (req, res) => {
    // req.pipe(request(`http://board:8080/reverse`)).pipe(res);
    // req.pipe(request(`${daprUrl}/boardapi/method/reverse`)).pipe(res);
    req.pipe(request(`${boardApi}/reverse`)).pipe(res);
})

app.post('/cp', async (req, res) => {
    // req.pipe(request(`http://cp:5000`)).pipe(res);
    // req.pipe(request(`${daprUrl}/cpapi/method`)).pipe(res);
    req.pipe(request(`${cpApi}/`)).pipe(res);
})



// Serve static files
app.use(express.static(path.join(__dirname, 'othello/build')));

app.get('*', function (_req, res) {
    res.sendFile(path.join(__dirname, 'othello/build', 'index.html'));
});

app.listen(process.env.PORT || port, () => console.log(`Listening on port ${port}!`));

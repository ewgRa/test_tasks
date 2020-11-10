const express = require('express');
const kafka = require('kafka-node');
const handlers = require('./src/handlers/vehicles');
const { log } = require('json-log');

const app = express();
app.use(express.json());

app.use(function (req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    res.header("Access-Control-Allow-Methods", "POST, GET, DELETE");
    next();
});

handlers.attach(app);

if(!module.parent) {
    app.listen(process.env.LISTEN_PORT, function () {
        log.info('App start listening', {port: process.env.LISTEN_PORT});
    });
}

module.exports = app;
const async = require('async');
const request = require('supertest');
const httpStatus = require('http-status-codes');
const app = require('../index');

describe('Unit testing the vehicle update location', function() {
    it('valid update', function(done) {
        async.series([
            (cb) => request(app)
                .post('/vehicles')
                .send({id: "test-location-update"})
                .expect(httpStatus.NO_CONTENT, cb),
            (cb) => request(app)
                .post('/vehicles/test-location-update/locations')
                .send({lat: 52.53, lng: 13.403, at: "2017-07-21T17:32:28Z"})
                .expect(httpStatus.NO_CONTENT, "", cb),
        ], done);
    });

    it('location outside city square', function(done) {
        async.series([
            (cb) => request(app)
                .post('/vehicles')
                .send({id: "test-location-update"})
                .expect(httpStatus.NO_CONTENT, cb),
            (cb) => request(app)
                .post('/vehicles/test-location-update/locations')
                .send({lat: 10, lng: 10, at: "2017-07-21T17:32:28Z"})
                .expect('X-AllygatorShuttle-Update-Ignored', 'Out of city square')
                .expect(httpStatus.NO_CONTENT, "", cb),
        ], done);
    });

    it('bad request update', function(done) {
        async.series([
            (cb) => request(app)
                .post('/vehicles')
                .send({id: "test-location-update"})
                .expect(httpStatus.NO_CONTENT, cb),
            (cb) => request(app)
                .post('/vehicles/test-location-update/locations')
                .send({lat: "not a number", lng: "not a number", at: "not a date"})
                .expect(httpStatus.BAD_REQUEST, `{"errors":[{"value":"not a number","msg":"'lat' should be a float","param":"lat","location":"body"},{"value":"not a number","msg":"'lng' should be a float","param":"lng","location":"body"},{"value":"not a date","msg":"'at' should be a datetime in ISO8601 format","param":"at","location":"body"}]}`, cb),
        ], done);
    });

    it('bad request required fields', function(done) {
        async.series([
            (cb) => request(app)
                .post('/vehicles')
                .send({id: "test-location-update"})
                .expect(httpStatus.NO_CONTENT, cb),
            (cb) => request(app)
                .post('/vehicles/test-location-update/locations')
                .send({})
                .expect(httpStatus.BAD_REQUEST, `{"errors":[{"msg":"'lat' should be a float","param":"lat","location":"body"},{"msg":"'lat' is required","param":"lat","location":"body"},{"msg":"'lng' should be a float","param":"lng","location":"body"},{"msg":"'lng' is required","param":"lng","location":"body"},{"msg":"'at' should be a datetime in ISO8601 format","param":"at","location":"body"},{"msg":"'at' is required","param":"at","location":"body"}]}`, cb),
        ], done);
    });

    it('vehicle not registered', function() {
        return request(app)
            .post('/vehicles/test-location-update-not-registered/locations')
            .send({lat: 10.2, lng: 12.2, at: "2017-07-21T17:32:28Z"})
            .expect(httpStatus.BAD_REQUEST, '{"errors":[{"msg":"Vehicle is not registered","param":"id","location":"path"}]}')
    });
});
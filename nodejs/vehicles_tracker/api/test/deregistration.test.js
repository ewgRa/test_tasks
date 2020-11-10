const async = require('async');
const request = require('supertest');
const httpStatus = require('http-status-codes');
const app = require('../index');

describe('Unit testing the vehicle de-registration', function() {
    const id = "test-de-registration";

    it('valid flow', function(done) {
        async.series([
            (cb) => request(app)
                .post('/vehicles')
                .send({id: id})
                .expect(httpStatus.NO_CONTENT, cb),
            (cb) => request(app)
                .delete('/vehicles/'+id)
                .send({})
                .expect(httpStatus.NO_CONTENT, "", cb),
            (cb) => request(app)
                .delete('/vehicles/'+id)
                .send({})
                .expect(httpStatus.BAD_REQUEST, cb),
        ], done);
    });

    it('not registered before', function() {
        return request(app)
            .delete('/vehicles/test-not-registered-de-registration')
            .send({})
            .expect(httpStatus.BAD_REQUEST, '{"errors":[{"msg":"Vehicle is not registered","param":"id","location":"path"}]}');
    });
});

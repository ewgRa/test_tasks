const async = require('async');
const request = require('supertest');
const httpStatus = require('http-status-codes');
const app = require('../index');

describe('Unit testing the vehicle registration', function() {

    it('valid registration', function() {
        return request(app)
            .post('/vehicles')
            .send({id: "valid"})
            .expect(httpStatus.NO_CONTENT, "")
    });

    it('id is not a string', function() {
        return request(app)
            .post('/vehicles')
            .send({id: 1})
            .expect(httpStatus.BAD_REQUEST, `{"errors":[{"value":1,"msg":"'id' should be a string","param":"id","location":"body"}]}`)
    });

    it('id missing', function() {
        return request(app)
            .post('/vehicles')
            .send({})
            .expect(httpStatus.BAD_REQUEST, `{"errors":[{"msg":"'id' should be a string","param":"id","location":"body"},{"msg":"'id' is required","param":"id","location":"body"}]}`)
    });
});

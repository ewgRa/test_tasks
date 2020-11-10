const assert = require('assert');
const async = require('async');
const request = require('supertest');
const httpStatus = require('http-status-codes');
const app = require('../index');
const { Client } = require('@elastic/elasticsearch');
const client = new Client({ node: process.env.ES_URL });
const storage = require('../src/storage');
const sleep = (waitTimeInMs) => new Promise(resolve => setTimeout(resolve, waitTimeInMs));

describe('Testing how records stored', function() {
    const id = "test-storage";

    it('whole flow', async function() {
        await request(app)
                .post('/vehicles')
                .send({id: id})
                .expect(httpStatus.NO_CONTENT);

        // Out of city square, should be ignored
        await request(app)
            .post('/vehicles/' + id + '/locations')
            .send({lat: 10.2, lng: 12.2, at: "2017-07-21T17:32:28Z"})
            .expect(httpStatus.NO_CONTENT);

        await request(app)
            .post('/vehicles/' + id + '/locations')
            .send({lat: 52.53, lng: 13.403, at: "2017-07-21T17:32:28Z"})
            .expect(httpStatus.NO_CONTENT);

        await waitRefresh();

        await request(app)
            .post('/vehicles/' + id + '/locations')
            .send({lat: 52.53, lng: 13.404, at: "2017-07-21T17:32:28Z"})
            .expect(httpStatus.NO_CONTENT);

        await waitRefresh();

        const { body } = await client.search({
            index: process.env.ES_INDEX,
            body: {
                query: {
                    match: { id: id }
                }
            }
        });

        assert.equal(body.hits.hits.length, 2);

        assert.deepEqual(body.hits.hits[1]._source, {
            id: 'test-storage',
            bearing: 89.99960316396367,
            type: 'location-update',
            lat: 52.53,
            lng: 13.404,
            at: '2017-07-21T17:32:28Z'
        });
    });
});

describe('Testing get latest update', function() {
    const id = "test-storage-latest-update";

    it('get latest update', async function() {
        await storage.save(id, "test", 0, 0, '2017-07-21T17:32:28Z', null);
        await storage.save(id, "test", 0, 0, '2019-07-21T17:32:28Z', null);
        await storage.save(id, "test", 0, 0, '2018-07-21T17:32:28Z', null);
        await waitRefresh();

        latest = await storage.latestUpdate(id);
        assert.equal(latest.at, '2019-07-21T17:32:28Z');
    });
});

async function waitRefresh() {
    await client.indices.refresh({ index: process.env.ES_INDEX });
}
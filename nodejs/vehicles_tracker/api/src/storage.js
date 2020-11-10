const { Client } = require('@elastic/elasticsearch');

const client = new Client({ node: process.env.ES_URL });

const registeredVehicles = {};
exports.registeredVehicles = registeredVehicles;

exports.save = async (id, recordType, lat, lng, at, bearing) => client.index({
  index: process.env.ES_INDEX,
  body: {
    id,
    type: recordType,
    lat,
    lng,
    bearing,
    at,
  },
});

exports.latestUpdate = async (id) => {
  const { body } = await client.search({
    index: process.env.ES_INDEX,
    sort: 'at:desc',
    size: 1,
    body: {
      query: {
        match: { id },
      },
    },
  });

  if (!body.hits.hits.length) {
    return null;
  }

  return body.hits.hits[0]['_source']; // eslint-disable-line dot-notation
};

const httpStatus = require('http-status-codes');
const { check, validationResult } = require('express-validator');
const kafka = require('../../kafka');
const storage = require('../../storage');
const navMath = require('../../nav_math');

const cityCenter = {
  lat: process.env.CITY_CENTER_LAT,
  lng: process.env.CITY_CENTER_LNG,
  radius: process.env.CITY_CENTER_RADIUS,
};

function calcBearing(latest, lat, lng) {
  if (latest == null) {
    return null;
  }

  return navMath.bearing(latest.lat, latest.lng, lat, lng);
}

function outOfCityCenter(req) {
  const centerDistance = navMath.distance(
    cityCenter.lat,
    cityCenter.lng,
    req.body.lat,
    req.body.lng,
  );

  return centerDistance > cityCenter.radius;
}

const validationChecks = [
  check('lat')
    .isFloat().withMessage("'lat' should be a float")
    .not()
    .isEmpty()
    .withMessage("'lat' is required"),
  check('lng')
    .isFloat().withMessage("'lng' should be a float")
    .not()
    .isEmpty()
    .withMessage("'lng' is required"),
  check('at')
    .isISO8601().withMessage("'at' should be a datetime in ISO8601 format")
    .not()
    .isEmpty()
    .withMessage("'at' is required"),
];

exports.attach = (app) => {
  app.post(
    '/vehicles/:id/locations',
    validationChecks,
    async (req, res) => {
      const errors = validationResult(req);
      if (!errors.isEmpty()) {
        return res.status(httpStatus.BAD_REQUEST).json({ errors: errors.array() });
      }

      if (!storage.registeredVehicles[req.params.id]) {
        return res.status(httpStatus.BAD_REQUEST)
          .json({ errors: [{ msg: 'Vehicle is not registered', param: 'id', location: 'path' }] });
      }

      if (outOfCityCenter(req)) {
        res.setHeader('X-AllygatorShuttle-Update-Ignored', 'Out of city square');
        return res.sendStatus(httpStatus.NO_CONTENT);
      }

      const bearing = calcBearing(
        await storage.latestUpdate(req.params.id),
        req.body.lat,
        req.body.lng,
      );

      await storage.save(req.params.id, 'location-update', req.body.lat, req.body.lng, req.body.at, bearing);

      kafka.sendMessage(JSON.stringify({
        type: 'location_update', id: req.params.id, lat: req.body.lat, lng: req.body.lng, bearing, at: req.body.at,
      }));

      return res.sendStatus(httpStatus.NO_CONTENT);
    },
  );
};

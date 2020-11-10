const httpStatus = require('http-status-codes');
const kafka = require('../../kafka');
const storage = require('../../storage');

exports.attach = (app) => {
  app.delete('/vehicles/:id',
    (req, res) => {
      if (!storage.registeredVehicles[req.params.id]) {
        return res.status(httpStatus.BAD_REQUEST)
          .json({ errors: [{ msg: 'Vehicle is not registered', param: 'id', location: 'path' }] });
      }

      kafka.sendMessage(JSON.stringify({ type: 'de-register', id: req.params.id }));
      delete storage.registeredVehicles[req.params.id];

      return res.sendStatus(httpStatus.NO_CONTENT);
    });
};

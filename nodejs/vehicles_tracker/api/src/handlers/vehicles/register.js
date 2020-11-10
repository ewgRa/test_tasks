const httpStatus = require('http-status-codes');
const { check, validationResult } = require('express-validator');
const kafka = require('../../kafka');
const storage = require('../../storage');

exports.attach = (app) => {
  app.post('/vehicles', [
    check('id')
      .isString().withMessage("'id' should be a string")
      .not()
      .isEmpty()
      .withMessage("'id' is required"),
  ],
  (req, res) => {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(httpStatus.BAD_REQUEST).json({ errors: errors.array() });
    }

    kafka.sendMessage(JSON.stringify({ type: 'register', id: req.body.id }));
    storage.registeredVehicles[req.body.id] = true;
    return res.sendStatus(httpStatus.NO_CONTENT);
  });
};

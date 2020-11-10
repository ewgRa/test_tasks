const register = require('./vehicles/register');
const deregister = require('./vehicles/deregister');
const updateLocation = require('./vehicles/update_location');
const broadcast = require('./vehicles/broadcast');

exports.attach = (app) => {
  register.attach(app);
  deregister.attach(app);
  updateLocation.attach(app);
  broadcast.attach(app);
};

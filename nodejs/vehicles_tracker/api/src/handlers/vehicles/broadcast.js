const expressWs = require('express-ws');
const kafka = require('../../kafka');

exports.attach = (app) => {
  const wsApp = expressWs(app);
  app.ws('/vehicles/broadcast', () => {});

  const aWss = wsApp.getWss('/vehicles/broadcast');

  kafka.consumeMessages((message) => {
    aWss.clients.forEach((client) => {
      client.send(message.value);
    });
  });
};

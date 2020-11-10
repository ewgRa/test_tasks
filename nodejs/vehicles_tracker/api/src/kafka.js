const kafka = require('kafka-node');
const { log } = require('json-log');

exports.sendMessage = (message) => {
  const client = new kafka.KafkaClient({ kafkaHost: process.env.KAFKA_HOST });
  const producer = new kafka.Producer(client);

  const payloads = [
    { topic: process.env.KAFKA_TOPIC, messages: message },
  ];

  producer.on('ready', () => {
    producer.send(payloads, (err) => {
      if (err) { log.error("Can't sent data to kafka", err); }
    });
  });

  producer.on('error', (err) => {
    log.error('Kafka error', err);
  });
};

exports.consumeMessages = (cb) => {
  const client = new kafka.KafkaClient({ kafkaHost: process.env.KAFKA_HOST });

  const consumer = new kafka.Consumer(
    client,
    [],
    { fromOffset: true },
  );

  consumer.on('message', cb);

  consumer.addTopics([
    { topic: process.env.KAFKA_TOPIC, partition: 0, offset: 0 },
  ], () => log.info('kafka topic added to consumer for listening', { topic: process.env.KAFKA_TOPIC }));
};

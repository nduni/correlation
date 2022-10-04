FROM rabbitmq:3.8-management-alpine

RUN rabbitmq-plugins enable --offline rabbitmq_amqp1_0

ADD rabbitmq.conf /etc/rabbitmq/
ADD definitions.json /etc/rabbitmq/
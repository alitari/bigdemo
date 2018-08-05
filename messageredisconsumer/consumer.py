#!/usr/bin/env python
import pika
import os

rabbitmq_host = os.environ['RABBITMQ_HOST']
queue_name = 'redisconsumer'
exchange_name = 'messages-exchange'

credentials = pika.PlainCredentials(username='user', password='user')
print(' using rabbitmq host \'% s\' with user \'% s\'' %
      (rabbitmq_host, credentials.username))

connection = pika.BlockingConnection(
    pika.ConnectionParameters(host=rabbitmq_host, credentials=credentials))
channel = connection.channel()

result = channel.queue_declare(queue=queue_name, durable=True)

channel.queue_bind(exchange=exchange_name,
                   queue=result.method.queue)


def message_received(ch, method, properties, body):
    print(" Received %r" % body)


print('consuming queue \'% s\' bound to exchange \'% s\' ...' %
      (queue_name, exchange_name))

channel.basic_consume(message_received,
                      queue=result.method.queue,
                      no_ack=True)

channel.start_consuming()

#!/usr/bin/env python
import pika
import redis
import os
import json


with open('english_words.txt') as f:
    english_words = f.readlines()
    english_words = [x.strip() for x in english_words]


gap_words = ['and', 'or', '.']
message_id_key = 'message_id'
rabbitmq_host = os.environ['RABBITMQ_HOST']
rabbitmq_user = os.environ['RABBITMQ_USER']
rabbitmq_password = os.environ['RABBITMQ_PASSWORD']

r = redis.Redis(host='bigdemo-redis-master', password='redis')

queue_name = os.environ['RABBITMQ_MESSAGE_QUEUE']
exchange_name = os.environ['RABBITMQ_EXCHANGE']

credentials = pika.PlainCredentials(
    username=rabbitmq_user, password=rabbitmq_password)
print(' using rabbitmq host \'% s\' with user \'% s\'' %
      (rabbitmq_host, credentials.username))

connection = pika.BlockingConnection(
    pika.ConnectionParameters(host=rabbitmq_host, credentials=credentials))
channel = connection.channel()

result = channel.queue_declare(queue=queue_name, durable=True)

channel.queue_bind(exchange=exchange_name,
                   queue=result.method.queue)


def message_received(ch, method, properties, body):

    m_dict = json.loads(body)

    m_id = store_message(m_dict)

    words_all = m_dict['text'].split()
    words_filtered = [w for w in words_all if w in english_words]

    for word in words_filtered:
        store_word(word, m_id)


def store_message(json):
    m_text = json['text']
    m_creationTime = json['creationTime']
    m_author = json['author']
    m_id = r.incr(message_id_key)
    print(" Received message text: %s" % m_text)
    if r.rpush(m_id, m_text,m_author,m_creationTime):
        print(" and stored with id %s" % m_id)
    return m_id


def store_word(word, m_id):
    if r.rpush(word, m_id):
        print(" stored word %s" % word)


print('consuming queue \'% s\' bound to exchange \'% s\' ...' %
      (queue_name, exchange_name))

channel.basic_consume(message_received,
                      queue=result.method.queue,
                      no_ack=True)

channel.start_consuming()

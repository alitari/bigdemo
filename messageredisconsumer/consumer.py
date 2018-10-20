#!/usr/bin/env python
import pika
import redis
import os
import json
import sys


class KVStorage:

    def __init__(self, redis_host, redis_password):
        self.r = redis.Redis(host=redis_host,
                             password=redis_password)

    def store_message(self, text, author, creationTime):
        print('storing values: { text: \'% s\' , author: \'% s\' , creationTime: \'% s\' }...' %
              (text, author, creationTime))
        m_id = self.r.incr('message_id')

        if self.r.rpush(m_id, text, author, creationTime):
            print("successfully with id %s" % m_id)
            self.r.incr('message_count')
        return m_id

    def store_word(self, word, key):
        print(' adding message_id: \'% s\' to word: \'% s\'...' %
              (key, word))
        if self.r.rpush(word, key):
            print(" successfully")


class MessageProducer:

    def __init__(self, host, user, password, exchange, queue, listener_method):
        self.rabbitmq_host = host
        self.rabbitmq_user = user
        self.rabbitmq_password = password
        self.exchange_name = exchange
        self.queue_name = queue
        self.listener_method = listener_method
        self.init_rabbitmq()

    def init_rabbitmq(self):
        self.credentials = pika.PlainCredentials(
            username=self.rabbitmq_user, password=self.rabbitmq_password)
        print(' using rabbitmq host \'% s\' with user \'% s\'' %
              (self.rabbitmq_host, self.credentials.username))
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(host=self.rabbitmq_host, credentials=self.credentials))
        self.channel = self.connection.channel()

        self.result = self.channel.queue_declare(
            queue=self.queue_name, durable=True)

        self.channel.exchange_declare(
            exchange=self.exchange_name, exchange_type='fanout', durable=True)

        self.channel.queue_bind(exchange=self.exchange_name,
                                queue=self.result.method.queue)

        self.channel.basic_consume(self.consume_method,
                                   queue=self.result.method.queue,
                                   no_ack=True)

        print('queue \'% s\' bound to exchange \'% s\'' %
              (self.queue_name, self.exchange_name))

    def consume_method(self, ch, method, properties, body):
        self.listener_method(body)

    def start_consuming(self):
        print('start consuming...')
        self.channel.start_consuming()


# with open('english_words.txt') as f:
#     english_words = f.readlines()
#     english_words = [x.strip() for x in english_words]


def message_received(body):

    try:
        m_dict = json.loads(body)

        m_text = m_dict['text']
        m_epochSecond = m_dict['creationTime']['epochSecond']
        m_nano = m_dict['creationTime']['nano']
        m_author = m_dict['author']
        m_id = ds.store_message(
            m_text, m_author, '{ "epochSecond": %d , "nano": %d }' % (m_epochSecond, m_nano))

        words_all = m_dict['text'].split()
        #words_filtered = [w for w in words_all]

        for word in words_all:
            ds.store_word(word, m_id)
    except:
        e0 = sys.exc_info()[0]
        e1 = sys.exc_info()[1]
        print(e0)
        print(e1)
        print('message body=\'% s\'' % body)


if __name__ == '__main__':
    producer = MessageProducer(os.environ['RABBITMQ_HOST'], os.environ['RABBITMQ_USER'], os.environ['RABBITMQ_PASSWORD'],
                               os.environ['RABBITMQ_EXCHANGE'], os.environ['RABBITMQ_MESSAGE_QUEUE'], message_received)
    ds = KVStorage(os.environ['REDIS_HOST'], os.environ['REDIS_PASSWORD'])
    producer.start_consuming()

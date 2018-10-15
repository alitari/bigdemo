import unittest
import consumer
import os

from mock import Mock, MagicMock, call


class ConsumerTest(unittest.TestCase):

    def test_valid_message_received_redis(self):
        consumer.ds = consumer.KVStorage('localhost', 'redis')
        consumer.message_received(
            '{ "text":"Hello", "author":"Alex", "creationTime":73364}')

    def test_valid_message_received_mock(self):
        consumer.ds = Mock(spec=consumer.KVStorage)
        consumer.ds.store_message.return_value = 11
        consumer.message_received(
            '{ "text":"Hello world", "author":"Alex", "creationTime":73364}')

        consumer.ds.store_message.assert_called_once_with(
            "Hello world", "Alex", 73364)

        self.assertEqual(consumer.ds.store_word.call_args_list, [
                         call(u'Hello', 11), call(u'world', 11)])

    def test_message_with_wrong_key_received_mock(self):
        consumer.ds = Mock(spec=consumer.KVStorage)
        consumer.message_received(
            '{ "wrongtext":"Hello world", "author":"Alex", "creationTime":73364}')

        self.assertFalse(consumer.ds.store_message.called)
        self.assertFalse(consumer.ds.store_word.called)

    def test_non_json_message_received_mock(self):
        consumer.ds = Mock(spec=consumer.KVStorage)
        consumer.message_received(
            'no json')

        self.assertFalse(consumer.ds.store_message.called)
        self.assertFalse(consumer.ds.store_word.called)


if __name__ == '__main__':
    unittest.main()

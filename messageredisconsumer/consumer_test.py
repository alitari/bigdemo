import unittest
import consumer
import os

from mock import Mock, MagicMock, call


class ConsumerTest(unittest.TestCase):

    @unittest.skip("can only run with redis")
    def test_redis_valid_message_received(self):
        consumer.ds = consumer.KVStorage('localhost', 'redis')
        consumer.message_received(
            '{ "text":"Hello", "author":"Alex", "creationTime":73364}')

    def test_mock_valid_message_received(self):
        consumer.ds = Mock(spec=consumer.KVStorage)
        consumer.ds.store_message.return_value = 11
        consumer.message_received(
            '{ "text":"Hello world", "author":"Alex", "creationTime":73364}')

        consumer.ds.store_message.assert_called_once_with(
            "Hello world", "Alex", 73364)

        self.assertEqual(consumer.ds.store_word.call_args_list, [
                         call(u'Hello', 11), call(u'world', 11)])

    def test_mock_message_with_wrong_key_received(self):
        consumer.ds = Mock(spec=consumer.KVStorage)
        consumer.message_received(
            '{ "wrongtext":"Hello world", "author":"Alex", "creationTime":73364}')

        self.assertFalse(consumer.ds.store_message.called)
        self.assertFalse(consumer.ds.store_word.called)

    def test_mock_non_json_message_received(self):
        consumer.ds = Mock(spec=consumer.KVStorage)
        consumer.message_received(
            'no json')

        self.assertFalse(consumer.ds.store_message.called)
        self.assertFalse(consumer.ds.store_word.called)


def run_suite():
    loader = unittest.TestLoader()
    loader.testMethodPrefix = "test_"
    suite = loader.discover('.', '*_test.py')
    unittest.TextTestRunner(verbosity=2).run(suite)


if __name__ == '__main__':
    run_suite()

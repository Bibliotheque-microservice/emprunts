import pika
import pytest

@pytest.fixture
def rabbitmq_connection():
    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    yield connection
    connection.close()

def test_rabbitmq_queue(rabbitmq_connection):
    channel = rabbitmq_connection.channel()
    channel.queue_declare(queue='test_queue')

    # Envoi d'un message
    channel.basic_publish(exchange='', routing_key='test_queue', body='Hello RabbitMQ!')

    # Vérification de la réception
    method_frame, header_frame, body = channel.basic_get(queue='test_queue', auto_ack=True)
    assert body == b'Hello RabbitMQ!'

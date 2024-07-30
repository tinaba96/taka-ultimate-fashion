import tensorflow as tf


def handler(event, context):
    return {
        'statusCode': 200,
        'body':
            {
                'predict': tf.__version__,
            }
    }

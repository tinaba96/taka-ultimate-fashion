import tensorflow as tf
import numpy as np
import boto3
import time


s3_client = boto3.client('s3')
BUCKET_NAME = 'tuf-products-data'


model = tf.keras.applications.EfficientNetB0(include_top=False, pooling="avg")

def image2vec(raw):
    # raw = tf.io.read_file(image_path)
    image = tf.image.decode_jpeg(raw, channels=3)
    image = tf.image.resize(image, [224, 224])
    vec = model.predict(np.array([image.numpy()]))[0]
    return vec

def cos_sim(v1, v2):
    return np.dot(v1, v2) / (np.linalg.norm(v1) * np.linalg.norm(v2))


def handler(event, context):
    OBJECT_KEY_NAME = 'products-images/product_image_'+'1128'+'.jpg'
    response = s3_client.get_object(Bucket=BUCKET_NAME, Key=OBJECT_KEY_NAME)
    body = response['Body'].read()
    # print(body)
    # img = image2vec(body)
    # s3 = boto3.resource('s3')
    # bucket = s3.Bucket(BUCKET_NAME)
    # time.sleep(10) 
    # obj = bucket.Object(OBJECT_KEY_NAME).get()
    # body = obj['Body'].read()
    img = image2vec(body)
    print(img)
    return {
        'statusCode': 200,
        'body':
            {
                # 'predict': tf.__version__,
                'product_id': [4, 33, 45]
            }
    }
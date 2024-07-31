import tensorflow as tf
import numpy as np
import boto3
import time
import json


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
    products = []
    response = boto3.client('lambda').invoke(
        FunctionName = 'test2',
        InvocationType='RequestResponse'
    )
    rds = json.loads(response['Payload'].read())
    time.sleep(5) 
    OBJECT_KEY_NAME = 'products-images/product_image_'+'1128'+'.jpg'
    response = s3_client.get_object(Bucket=BUCKET_NAME, Key=OBJECT_KEY_NAME)
    body = response['Body'].read()
    img = image2vec(body)
    for i in rds['body']:
        products.append(i[0])
        OBJECT_KEY_NAME2 = 'products-images/product_image_'+str(i)+'.jpg'
        response2 = s3_client.get_object(Bucket=BUCKET_NAME, Key=OBJECT_KEY_NAME2)
        body2 = response2['Body'].read()
        img2 = image2vec(body2)
        diff = cos_sim(img, img2)
        print(diff)
    
        # print(img)
        # print(rds)
    return {
        'statusCode': 200,
        'body':
            {
                # 'predict': tf.__version__,
                'product_id': [4, 33, 45]
            }
    }
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
    start_time = time.time()
    products = []
    uniq = set()
    response = boto3.client('lambda').invoke(
        FunctionName = 'test2',
        InvocationType='RequestResponse'
    )
    rds = json.loads(response['Payload'].read())
    time.sleep(2) 
    # print(rds['body']['all'])
    for i in rds['body']['saves']:
        OBJECT_KEY_NAME = 'products-images/product_image_'+str(i[0])+'.jpg'
        response = s3_client.get_object(Bucket=BUCKET_NAME, Key=OBJECT_KEY_NAME)
        body = response['Body'].read()
        img = image2vec(body)
        # print(rds['body']['all'][str(i[1])])
        for k in rds['body']['all'][str(i[1])]:
            if time.time() - start_time > 10:
                return {
                    'statusCode': 200,
                    'headers': {
                        'Access-Control-Allow-Headers': 'Content-Type',
                        'Access-Control-Allow-Origin': '*',
                        'Access-Control-Allow-Methods': 'OPTIONS,POST,GET'
                    },
                    'body':
                        {
                            'products': list(products)
                        }
                }
            OBJECT_KEY_NAME2 = 'products-images/product_image_'+str(k[0])+'.jpg'
            response2 = s3_client.get_object(Bucket=BUCKET_NAME, Key=OBJECT_KEY_NAME2)
            body2 = response2['Body'].read()
            img2 = image2vec(body2)
            diff = cos_sim(img, img2)
            # print(diff)
            if diff > 0.85 and i[0] != k[0] and k[0] not in uniq:
                uniq.add(k[0])
                item = {
                    'ID': k[0],
                    'Name': k[1],
                    'ImageURL': k[2],
                    'Price': k[3]
                }
                products.append(item)

    
    return {
        'statusCode': 200,
        'headers': {
            'Access-Control-Allow-Headers': 'Content-Type',
            'Access-Control-Allow-Origin': 'https://www.example.com',
            'Access-Control-Allow-Methods': 'OPTIONS,POST,GET'
        },
        'body':
            {
                'products': products
            }
    }
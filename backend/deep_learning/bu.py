import json
# import os
import boto3
import numpy as np
import tensorflow as tf
# import pymysql
import time

# host = os.environ['RDS_HOST_NAME']
# user = os.environ['USER']
# password = os.environ['PASS']
# db = os.environ['DB']

# connection = pymysql.connect(host=host, user=user, password=password, database=db)

model = tf.keras.applications.EfficientNetB0(include_top=False, pooling="avg")
BUCKET_NAME = 'tuf-products-data'

# def image2vec(image_path):
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
    # 別Lambda関数の呼び出し
    # response = boto3.client('lambda').invoke(
    #     FunctionName = 'test2',
    #     InvocationType='RequestResponse'
    # )
    # rdb = json.loads(response['Payload'].read())

    # time.sleep(10) 
    # for i in rdb['body']:
    #     products.append(i[0])

    # with connection.cursor() as cursors:
    #     cursors.execute('SELECT product_id FROM saves')
    #     saves_productID = cursors.fetchall()
    # print(saves_productID)
    # for product in saves_productID:
    #     print(product[0])

    OBJECT_KEY_NAME = 'products-images/product_image_'+'1128'+'.jpg'
    # s3_client = boto3.client('s3')
    # response = s3_client.get_object(Bucket=BUCKET_NAME, Key=OBJECT_KEY_NAME)
    # body = response['Body'].read()
    # response2 = s3_client.get_object(Bucket=BUCKET_NAME, Key=OBJECT_KEY_NAME2)
    s3 = boto3.resource('s3')
    bucket = s3.Bucket(BUCKET_NAME)
    obj = bucket.Object(OBJECT_KEY_NAME).get()
    # print(obj['body'].read())
    body = obj['Body'].read()
    # print(body)


    img = image2vec(body)
    print(img)
    return {
        'statusCode': 200,
        'body':
            {
                'predict': tf.__version__,
                'product_id': [4, 33, 45]
            }
    }

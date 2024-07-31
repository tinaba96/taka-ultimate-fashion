try:
    import unzip_requirements
except ImportError:
    pass

import io
import os
from boto3.session import Session
from tensorflow.keras.models import load_model
from tensorflow.keras.preprocessing import image
from tensorflow.keras.applications.resnet50 \
    import preprocess_input, decode_predictions
import numpy as np

MODEL_BUCKET ='image-classification-lambda'
MODEL_NAME = 'model_resnet50_imagenet.h5'
MODEL_PATH = os.path.join('/tmp', MODEL_NAME)

# (2) モデルのダウンロード
session = Session()
s3_client = session.client('s3')
s3_client.download_file(MODEL_BUCKET, MODEL_NAME, MODEL_PATH)

model = None


def classify(event, context):
    global model
    if model is None:
        model = load_model(MODEL_PATH, compile=False)

    # (3) 画像のダウンロード
    s3_object = s3_client.get_object(
        Bucket=event['bucket'], Key=event['filename'])
    image_data = io.BytesIO(s3_object['Body'].read())
    img = image.load_img(image_data, target_size=(224, 224))
    x = image.img_to_array(img)
    x = np.expand_dims(x, axis=0)
    x = preprocess_input(x)

    # (4) 推論
    preds = model.predict(x)
    results = decode_predictions(preds, top=3)[0]

    return [{'class': r[1], 'score': float(r[2])} for r in results]
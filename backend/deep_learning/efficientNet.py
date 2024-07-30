import numpy as np
import tensorflow as tf

model = tf.keras.applications.EfficientNetB0(include_top=False, pooling="avg")

def image2vec(image_path):
    raw = tf.io.read_file(image_path)
    image = tf.image.decode_jpeg(raw, channels=3)
    image = tf.image.resize(image, [224, 224])
    vec = model.predict(np.array([image.numpy()]))[0]
    return vec

print(image2vec('../image1.webp'))

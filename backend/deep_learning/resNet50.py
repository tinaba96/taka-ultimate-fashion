from tensorflow.keras.applications.resnet50 import ResNet50

model = ResNet50(weights='imagenet')
model.save('/path/to/model_resnet50_imagenet.h5', 
           include_optimizer=False)
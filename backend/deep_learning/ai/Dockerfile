# FROM public.ecr.aws/lambda/python:3.8
# FROM python:3.8-slim


FROM --platform=linux/amd64 public.ecr.aws/lambda/python:3.8

# COPY . ${LAMBDA_TASK_ROOT}
# RUN pip install --upgrade pip && pip install versioned-hdf5 && pip install h5py && pip install --upgrade pip setuptools wheel && pip install -r requirements.txt
# CMD ["app.handler"]

# FROM public.ecr.aws/lambda/python:3.6.2022.01.11.19
WORKDIR ${LAMBDA_TASK_ROOT}

# RUN yum update -y

# COPY src/ .
COPY . ${LAMBDA_TASK_ROOT}
COPY requirements.txt requirements.txt

RUN pip3 install --upgrade pip
RUN pip3 install -r requirements.txt

# CMD ["lambda_function.lambda_handler"]
CMD ["app.handler"]
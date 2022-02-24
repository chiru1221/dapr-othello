# FROM python:3.7
# EXPOSE 5000
# WORKDIR /root/src
# COPY ./cp/requirements.txt /root/src/requirements.txt
# RUN pip install -r requirements.txt

FROM pytorch/pytorch:latest
WORKDIR /workspace
COPY ./cp/alpha-zero-general/utils.py /workspace/utils.py
COPY ./cp/alpha-zero-general/alpha_zero.py /workspace/alpha_zero.py
COPY ./cp/alpha-zero-general/MCTS.py /workspace/MCTS.py
COPY ./cp/alpha-zero-general/Game.py /workspace/Game.py
COPY ./cp/alpha-zero-general/NeuralNet.py /workspace/NeuralNet.py
COPY ./cp/alpha-zero-general/pretrained_models/othello/pytorch/8x8_100checkpoints_best.pth.tar \
/workspace/8x8_100checkpoints_best.pth.tar
COPY ./cp/alpha-zero-general/othello /workspace/othello
RUN pip install flask flask_cors pytest grpcio grpcio-tools
# RUN apt update
# RUN apt install -y protobuf-compiler

ENV PYTHON_ENV=development
EXPOSE 5000

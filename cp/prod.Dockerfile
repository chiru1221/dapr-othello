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
COPY ./cp/main.py /workspace/main.py
COPY ./cp/board_pb2.py /workspace/board_pb2.py
COPY ./cp/board_pb2_grpc.py /workspace/board_pb2_grpc.py
COPY ./cp/cp_pb2.py /workspace/cp_pb2.py
COPY ./cp/cp_pb2_grpc.py /workspace/cp_pb2_grpc.py
COPY ./proto /workspace/proto

RUN pip install grpcio grpcio-tools
ENV PYTHON_ENV=production
EXPOSE 5000
ENTRYPOINT [ "python", "main.py" ]

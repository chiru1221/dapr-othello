FROM pytorch/pytorch:latest
WORKDIR /workspace
COPY utils.py /workspace/utils.py
COPY main.py /workspace/main.py
COPY MCTS.py /workspace/MCTS.py
COPY Game.py /workspace/Game.py
COPY NeuralNet.py /workspace/NeuralNet.py
COPY pretrained_models/othello/pytorch/8x8_100checkpoints_best.pth.tar \
/workspace/8x8_100checkpoints_best.pth.tar
COPY othello /workspace/othello

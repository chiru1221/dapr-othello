# Ref: https://qiita.com/bpzAkiyama/items/7d50f0e1ef1e262df984
from MCTS import MCTS as mcts
import numpy as np
from othello.OthelloGame import OthelloGame as Game
from othello.pytorch.NNet import NNetWrapper as nn
from utils import dotdict

MODEL = nn(Game(8))
MODEL.load_checkpoint('./', '8x8_100checkpoints_best.pth.tar')
MCTS = mcts(Game(8), MODEL, dotdict({'numMCTSSims': 25, 'cpuct': 1.0}))

'''
board: np.ndarray
black: -1
white: 1
space: 0

args: board
if stone is black
-> board * -1
else
-> board

MODEL put white(1) stone
-1 1
1  -1
-->>
1 -1 1
   1 -1
'''

def predict(board):
    action = np.argmax(MCTS.getActionProb(board, temp=1.0))
    return int(action // 8), int(action % 8)

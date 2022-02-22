# import logging

# import coloredlogs

# from Coach import Coach
from MCTS import MCTS
import numpy as np
from othello.OthelloGame import OthelloGame as Game
from othello.pytorch.NNet import NNetWrapper as nn
from utils import dotdict

# log = logging.getLogger(__name__)

# coloredlogs.install(level='INFO')  # Change this to DEBUG to see more info.

# args = dotdict({
#     'numIters': 1000,
#     'numEps': 100,              # Number of complete self-play games to simulate during a new iteration.
#     'tempThreshold': 15,        #
#     'updateThreshold': 0.6,     # During arena playoff, new neural net will be accepted if threshold or more of games are won.
#     'maxlenOfQueue': 200000,    # Number of game examples to train the neural networks.
#     'numMCTSSims': 25,          # Number of games moves for MCTS to simulate.
#     'arenaCompare': 40,         # Number of games to play during arena play to determine if new net will be accepted.
#     'cpuct': 1,

#     'checkpoint': './temp/',
#     'load_model': True,
#     'load_folder_file': ('/workspace/pretrained_models/othello/pytorch/','8x8_100checkpoints_best.pth.tar'),
#     'numItersForTrainExamplesHistory': 20,

# })


def main():
    # log.info('Loading %s...', Game.__name__)
    g = Game(8)

    # log.info('Loading %s...', nn.__name__)
    nnet = nn(g)

    # if args.load_model:
    # log.info('Loading checkpoint "%s/%s"...', args.load_folder_file[0], args.load_folder_file[1])
    nnet.load_checkpoint('./', '8x8_100checkpoints_best.pth.tar')
    # log.info('Loading the Coach...')
    # c = Coach(g, nnet, args)
    args = dotdict({'numMCTSSims': 25, 'cpuct': 1.0})
    mcts = MCTS(g, nnet, args)
    # return np.array, -> self impl ok
    board = g.getInitBoard()
    print(board)
    ' black -> -1, white -> 1, none -> 0'
    # curPlayer represent color, not necessary change
    curPlayer = 1
    # getCanonicalForm is not neccesary
    action = np.argmax(mcts.getActionProb(g.getCanonicalForm(board, curPlayer), temp=1.0))
    print(action)
    # board, curPlayer = g.getNextState(board, curPlayer, action)
    # print(board, curPlayer)
    # if args.load_model:
    #     log.info("Loading 'trainExamples' from file...")
    #     c.loadTrainExamples()

    # log.info('Starting the learning process 🎉')
    # c.learn()
# https://qiita.com/bpzAkiyama/items/7d50f0e1ef1e262df984

if __name__ == "__main__":
    main()

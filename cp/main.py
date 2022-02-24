import grpc
import cp_pb2
import cp_pb2_grpc
import board_pb2
import board_pb2_grpc
from concurrent import futures
from alpha_zero import predict
import re
import os
import numpy as np
import random

# create files related grpc
# python -m grpc_tools.protoc -I./proto --python_out=. --grpc_python_out=. ./proto/board.proto
# python -m grpc_tools.protoc -I./proto --python_out=. --grpc_python_out=. ./proto/cp.proto

ERROR_RES = lambda err: jsonify({'error': err})
NOPUT_RES = lambda: jsonify({'x': -1, 'y': -1})
STONE_STR_TO_INT = {'w': 1, 'b': -1}

DAPR_PORT = os.getenv('DAPR_GRPC_PORT', 50001)
BOARDAPI = 'localhost:{}'.format(DAPR_PORT) \
                if os.getenv('PYTHON_ENV') == 'production' else 'board:8080'
METADATA = (('dapr-app-id', 'boardapi'),)
    

class CpApi(cp_pb2_grpc.CpApiServicer):
    def Attack(self, request, context):
        level = request.level
        x = -1
        y = -1
        # checkputable
        squares = putable(board_pb2.Board(
            stone=request.stone,
            squares=request.squares,
        ))
        if 'p' in squares:
            request.squares = squares
            if level == 1:
                x, y = self.cp_lv1(request)
            elif level == 2:
                x, y = self.cp_lv2(request)
            elif level == 3:
                x, y = self.cp_lv3(request)
        return cp_pb2.Res(x=x, y=y)
    
    def cp_lv1(self, request):
        indexes = [m.span()[0] for m in re.finditer('p', request.squares)]
        put_idx = random.choice(indexes)
        x, y = int(put_idx // 8), int(put_idx % 8)
        return x, y
    
    def cp_lv2(self, request):
        indexes = [m.span()[0] for m in re.finditer('p', request.squares)]
        results = list()
        for put_idx in indexes:
            results.append(
                (
                    reverse(board_pb2.Board(
                        stone=request.stone,
                        x=int(put_idx // 8),
                        y=int(put_idx % 8),
                        squares=request.squares,
                    )),
                    int(put_idx // 8),
                    int(put_idx % 8),
                )
            )
        scores = np.array([[calc_score(result[0], request.stone), result[1], result[2]] for result in results])
        xy = scores[np.argmax(scores[:, 0])][1:]
        return xy[0], xy[1]
    
    def cp_lv3(self, request):
        board = string_to_ndarr(request.squares, 'w', 1) + string_to_ndarr(request.squares, 'b', -1)
        xy = predict(board*STONE_STR_TO_INT[request.stone])
        return xy[0], xy[1]


def string_to_ndarr(squares, str_stone, int_stone):
    arr_squares = np.zeros((8, 8))
    indexes = [m.span()[0] for m in re.finditer(str_stone, squares)]
    for idx in indexes:
        arr_squares[int(idx // 8)][int(idx % 8)] = int_stone
    return arr_squares

def calc_score(squares, stone):
    return len(re.findall(stone, squares))

def reverse(board):
    with grpc.insecure_channel(BOARDAPI) as channel:
        stub = board_pb2_grpc.BoardApiStub(channel)
        response = stub.Reverse(request=board, metadata=METADATA)
    return response.squares

def putable(board):
    with grpc.insecure_channel(BOARDAPI) as channel:
        stub = board_pb2_grpc.BoardApiStub(channel)
        response = stub.Putable(request=board, metadata=METADATA)
    return response.squares

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cp_pb2_grpc.add_CpApiServicer_to_server(CpApi(), server)
    server.add_insecure_port('[::]:5000')
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    print("start server")
    serve()

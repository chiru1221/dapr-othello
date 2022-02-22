import flask
from flask import request, jsonify
from flask_cors import CORS
import requests
import math
import sys
import json
import re
import os
import random
from multiprocessing import Pool
from functools import partial
import numpy as np
from alpha_zero import predict

app = flask.Flask(__name__)
CORS(app)

ERROR_RES = lambda err: jsonify({'error': err})
NOPUT_RES = lambda: jsonify({'x': -1, 'y': -1})
STONE_STR_TO_INT = {'w': 1, 'b': -1}

DAPR_PORT = os.getenv('DAPR_HTTP_PORT', 3500)
BOARDAPI = 'http://localhost:{}/v1.0/invoke/boardapi/method'.format(DAPR_PORT) \
                if os.getenv('PYTHON_ENV') == 'production' else 'http://board:8080'


@app.route('/', methods=['POST'])
def router():
    req = request.json
    if req is None:
        return ERROR_RES('no json request'), 400
    # Validate request
    if 'level' not in req:
        return ERROR_RES('not found `level`'), 400
    elif req['level'] < 1 or req['level'] > 3:
        return ERROR_RES('invalid `level` value'), 400

    if 'stone' not in req:
        return ERROR_RES('not found `stone`'), 400
    elif req['stone'] != 'w' and req['stone'] != 'b':
        return ERROR_RES('invalid `stone` value'), 400

    if 'squares' not in req:
        return ERROR_RES('not found `squares`'), 400
    
    # Get putable points
    r = request_to_putable(req['stone'], req['squares'])
    if r.status_code != 200:
        return ERROR_RES('putable server error'), 500
    elif 'p' not in r.json()['squares']:
        return NOPUT_RES()
    
    if req['level'] == 1:
        return cp_lv1(req['stone'], req['squares'], r)
    elif req['level'] == 2:
        return cp_lv2(req['stone'], req['squares'], r)
    elif req['level'] == 3:
        return cp_lv3(req['stone'], req['squares'], r)
    return ERROR_RES('internal server'), 500

# random select
def cp_lv1(stone, squares, r):
    indexes = [m.span()[0] for m in re.finditer('p', r.json()['squares'])]
    put_idx = random.choice(indexes)
    x, y = int(put_idx // 8), int(put_idx % 8)
    res = {'x': x, 'y': y}
    return jsonify(res)

def cp_lv2(stone, squares, r):
    indexes = [m.span()[0] for m in re.finditer('p', r.json()['squares'])]
    scores = list()
    with Pool(4) as pool:
        results = list()
        for put_idx in indexes:
            results.append(
                (
                    pool.apply_async(request_to_reverse, (stone, int(put_idx // 8), int(put_idx % 8), squares)),
                    int(put_idx // 8),
                    int(put_idx % 8),
                )
            )
        scores = [[calc_score(result[0].get(timeout=1).json()['squares'], stone), result[1], result[2]] for result in results]

    max_score = 0
    xy = [-1, -1]
    for score in scores:
        if max_score < score[0]:
            max_score = score[0]
            xy = score[1:]

    return jsonify({'x': xy[0], 'y': xy[1]})

def cp_lv3(stone, squares, r):
    board = string_to_ndarr(squares, 'w', 1) + string_to_ndarr(squares, 'b', -1)
    xy = predict(board*STONE_STR_TO_INT[stone])
    return jsonify({'x': xy[0], 'y': xy[1]})

def request_to_putable(stone, squares):
    # url = 'http://board:8080/putable'
    # daprPort = os.getenv("DAPR_HTTP_PORT", 3500)
    # url = 'http://localhost:{}/v1.0/invoke/boardapi/method/putable'.format(daprPort)
    payload = {'stone': stone, 'squares': squares}
    headers = {'content-type': 'application/json'}

    r = requests.post('{}/putable'.format(BOARDAPI), data=json.dumps(payload), headers=headers)
    return r

def request_to_reverse(stone, x, y, squares):
    # url = 'http://board:8080/reverse'
    # daprPort = os.getenv("DAPR_HTTP_PORT", 3500)
    # url = 'http://localhost:{}/v1.0/invoke/boardapi/method/reverse'.format(daprPort)
    payload = {'stone': stone, 'x': x, 'y': y, 'squares': squares}
    headers = {'content-type': 'application/json'}

    r = requests.post('{}/reverse'.format(BOARDAPI), data=json.dumps(payload), headers=headers)
    return r

def calc_score(squares, stone):
    return len(re.findall(stone, squares))

def string_to_ndarr(squares, str_stone, int_stone):
    arr_squares = np.zeros((8, 8))
    indexes = [m.span()[0] for m in re.finditer(str_stone, squares)]
    for idx in indexes:
        arr_squares[int(idx // 8)][int(idx % 8)] = int_stone
    return arr_squares

if __name__ == '__main__':
    print('start server')
    app.run(host='0.0.0.0')

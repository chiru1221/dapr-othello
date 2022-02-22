import pytest
# from main import app, router, request_to_putable, calc_score
from main import *
import requests


class Router_expected:
    def __init__(self, status_code, body):
        self.status_code = status_code
        self.body = body


class Response_mock:
    def __init__(self, status_code, body=None):
        self.status_code = status_code
        self.body = body
    
    def json(self):
        return self.body


@pytest.fixture()
def flask_app():
    app.config.update({
        "TESTING": True,
    })
    yield app

@pytest.fixture()
def client(flask_app):
    return flask_app.test_client()

@pytest.fixture()
def runner(flask_app):
    return flask_app.test_cli_runner()

@pytest.mark.parametrize(('data', 'res_json_key', 'putable_res', 'expected'), [
    (None, 'error', None, Router_expected(400, 'no json request')),
    ({}, 'error', None, Router_expected(400, 'not found `level`')),
    ({'level': 1}, 'error', None, Router_expected(400, 'not found `stone`')),
    ({'level': 1, 'stone': 'b'}, 'error', None, Router_expected(400, 'not found `squares`')),
    ({'level': 1, 'stone': 'b', 'squares': 'wbwb'}, 'error', 
        Response_mock(500), Router_expected(500, 'putable server error')),
    ({'level': 1, 'stone': 'b', 'squares': 'wbwb'}, 'x', 
        Response_mock(200, {'squares': 'bw'}), Router_expected(200, -1)),
])
def test_router(mocker, client, data, res_json_key, putable_res, expected):
    mocker.patch('main.request_to_putable', return_value=putable_res)
    res = client.post('/', json=data)
    assert res.status_code == expected.status_code
    assert res.json[res_json_key] == expected.body

@pytest.mark.parametrize(('squares', 'stone','expected'), [
    ('bbb', 'b', 3),
    ('www', 'b', 0),
    ('bwb', 'b', 2),
])
def test_calc_socre(squares, stone, expected):
    assert calc_score(squares, stone) == expected
    assert calc_score(squares, stone) == expected
    assert calc_score(squares, stone) == expected

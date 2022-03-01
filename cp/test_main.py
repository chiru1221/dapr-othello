import pytest
import cp_pb2
import cp_pb2_grpc
from main import *

@pytest.mark.parametrize(('level', 'putable_return', 'expected'), [
    (1, "bwwb", (-1, -1)),
    (1, "pppp", (1, 1)),
    (2, "pppp", (2, 2)),
    (3, "pppp", (3, 3)),
])
def test_attack(mocker, level, putable_return, expected):
    mocker.patch('main.putable', return_value=putable_return)
    mocker.patch('main.CpApi.cp_lv1', return_value=(1, 1))
    mocker.patch('main.CpApi.cp_lv2', return_value=(2, 2))
    mocker.patch('main.CpApi.cp_lv3', return_value=(3, 3))
    cp = CpApi()
    request = cp_pb2.Cp(level=level)
    assert cp.Attack(request, None) == cp_pb2.Res(x=expected[0], y=expected[1])

@pytest.mark.parametrize(('squares', 'stone','expected'), [
    ('bbb', 'b', 3),
    ('www', 'b', 0),
    ('bwb', 'b', 2),
])
def test_calc_socre(squares, stone, expected):
    assert calc_score(squares, stone) == expected
    assert calc_score(squares, stone) == expected
    assert calc_score(squares, stone) == expected
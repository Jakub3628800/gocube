import pytest
from cube import Cube


def test_init_cube():
    c = Cube()
    assert c.corner_permutation == list(range(0, 8))
    assert c.corner_orientation == [0] * 8
    assert c.edge_orientation == [0] * 12
    assert c.edge_permutation == list(range(0, 12))

def test_str():
    c = Cube()
    print(c)
    to_string = str(c)
    print(to_string)

from .utils import sum, reverse_string

def test_sum():
    assert sum(1, 2) == 3, "Should be 3"
    assert sum(5, 5) == 10, "Should be 10"
    assert sum(-5, 5) == 0, "Should be 0"
    assert sum(0, 0) == 0, "Should be 0"

def test_reverse_string():
    assert reverse_string('hello') == 'olleh', "Should be 'olleh'"
    assert reverse_string('World') == 'dlroW', "Should be 'dlroW'"
    assert reverse_string('Python') == 'nohtyP', "Should be 'nohtyP'"
    assert reverse_string('') == '', "Should be ''"
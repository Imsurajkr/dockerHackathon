import unittest
from calculator import calculator

# Test cases for the program
class TestCalculator(unittest.TestCase):

    def test_sum(self):
        total = calculator.sum_two_num(2,3)
        self.assertEqual(total, 5)

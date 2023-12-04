import unittest
from main import part1, create_matrix

class TestDay3(unittest.TestCase):
    def test_part1(self):
        matrix = create_matrix("test_input.txt")

        self.assertEqual(part1(matrix), 4361)

if __name__ == "__main__":
    unittest.main()

import unittest
from main import part1, create_matrix, part2

class TestDay3(unittest.TestCase):
    def test_part1(self):
        matrix = create_matrix("test_input.txt")

        self.assertEqual(part1(matrix), 4361)

    def test_part2(self):
        matrix = create_matrix("test_input.txt")
        self.assertEqual(part2(matrix), 467835)

if __name__ == "__main__":
    unittest.main()

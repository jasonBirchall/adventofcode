import unittest
from main import part1, part2

class Day4(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(part1("test_input.txt"), 13)

    def test_part2(self):
        self.assertEqual(part2("test_input.txt"), 30)

if __name__ == '__main__':
    unittest.main()

import unittest
from main import part1, part2

class TestDay2(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(part1("test_input_part1.txt"), 8)

    def test_part2(self):
        self.assertEqual(part2("test_input_part1.txt"), 2286)

if __name__ == '__main__':
    unittest.main()

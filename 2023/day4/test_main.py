import unittest
from main import part1

class Day4(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(part1("test_input.txt"), 13)

if __name__ == '__main__':
    unittest.main()

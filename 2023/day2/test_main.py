import unittest
from main import part1

class TestPart1(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(part1("test_input_part1.txt"), 8)

if __name__ == '__main__':
    unittest.main()

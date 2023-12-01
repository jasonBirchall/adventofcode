import unittest
from main import part1, part2

class TestDay1(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(part1("test_data.txt"), 142)
        
    def test_part2(self):
        self.assertEqual(part2("test_data_part2.txt"), 281)

if __name__ == '__main__':
    unittest.main()

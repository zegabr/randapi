import unittest
from src.app import get_python_string

class TestApp(unittest.TestCase):
    def test_get_python(self):
        response = get_python_string()
        self.assertEqual(response, 'hello from python')

if __name__ == '__main__':
    unittest.main()

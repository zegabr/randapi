import unittest
import threading
import time
from src.app import app  # Import the Flask app object from your actual app file
import requests

class TestApp(unittest.TestCase):
    def setUp(self):
        self.app = app.test_client()
        self.server_thread = threading.Thread(target=app.run, kwargs={"host": "127.0.0.1", "port": 5000})
        self.server_thread.start()
        time.sleep(2)  # Give the server some time to start

    def tearDown(self):
        self.server_thread.join()

    def test_get_data(self):
        print("opa")
        response = requests.get("http://127.0.0.1:5000/")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.text, "hello from python")

if __name__ == '__main__':
    unittest.main()

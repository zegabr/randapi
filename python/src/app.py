from flask import Flask

app = Flask(__name__)

@app.route("/",  methods=['GET'])
def get_python():
    return "hello from python"

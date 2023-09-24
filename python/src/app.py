from flask import Flask

app = Flask(__name__)

def get_python_string():
    # TODO: connect with database
    return "hello from python"

@app.route("/",  methods=['GET'])
def get_python():
    return get_python_string()

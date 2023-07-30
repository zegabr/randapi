FROM python:3.10-alpine

COPY . .

WORKDIR /src

RUN pip install -r requirements.txt
RUN export FLASK_APP=app.py

CMD [ "flask", "run", "--host=0.0.0.0"]

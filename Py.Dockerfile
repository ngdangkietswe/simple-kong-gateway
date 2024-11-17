FROM python:3.9-slim-bullseye

WORKDIR /app

COPY backend-servers/requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY backend-servers .

ENV FLASK_APP=main.py

CMD [ "python3", "-m", "flask", "run", "--host=0.0.0.0", "--port=5000" ]
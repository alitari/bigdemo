FROM python:2.7.15-slim-jessie

ENV PYTHONUNBUFFERED=0
WORKDIR /app
ADD requirements.txt .
ADD english_words.txt .
RUN pip install -r requirements.txt
ADD consumer.py .

CMD [ "python", "-u","./consumer.py" ]
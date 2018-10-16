FROM python:2.7.15-slim-jessie AS requirements
WORKDIR /app
ADD requirements.txt .
RUN pip install -r requirements.txt

FROM python:2.7.15-slim-jessie AS tests
ENV PYTHONUNBUFFERED=1
COPY --from=requirements /usr/local/lib/python2.7/site-packages /usr/local/lib/python2.7/site-packages
WORKDIR /app
ADD consumer.py .
ADD consumer_test.py .
RUN python consumer_test.py

FROM python:2.7.15-slim-jessie
COPY --from=tests /usr/local/lib/python2.7/site-packages /usr/local/lib/python2.7/site-packages
WORKDIR /app
ADD consumer.py .
CMD [ "python", "-u","./consumer.py" ]
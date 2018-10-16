FROM python:2.7.15-slim-jessie AS requirements
WORKDIR /app
ADD requirements.txt .
RUN pip install -r requirements.txt

FROM requirements AS tests
ENV PYTHONUNBUFFERED=1
ADD consumer.py .
ADD consumer_test.py .
RUN python -m unittest consumer_test

FROM requirements
COPY --from=tests /app/consumer.py consumer.py
CMD [ "python", "-u","./consumer.py" ]
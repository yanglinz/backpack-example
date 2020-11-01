# Container image that runs your code
FROM python:3.9.0-alpine3.12

RUN apk update && apk upgrade && apk add bash
RUN pip install pyyaml

RUN mkdir -p /app
COPY entrypoint.sh /app/entrypoint.sh
COPY main.py /app/main.py

ENTRYPOINT ["/app/entrypoint.sh"]

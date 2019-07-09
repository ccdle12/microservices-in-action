FROM python:3.7-slim-stretch

WORKDIR /market_service
COPY ./ ./

RUN pip install -r requirements.txt

CMD ["python", "run.py"]

FROM python:latest

WORKDIR /api_gateway
COPY ./ ./

RUN pip install -r requirements.txt

CMD ["python", "run.py"]

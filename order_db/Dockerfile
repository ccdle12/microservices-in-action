FROM postgres:latest

ENV POSTGRES_DB=$POSTGRES_DB
ENV POSTGRES_USER=$POSTGRES_USER
ENV POSTGRES_PASSWORD=$POSTGRES_PASSWORD

# COPY psql_dump.sql /docker-entrypoint-initdb.d
COPY ./ /docker-entrypoint-initdb.d

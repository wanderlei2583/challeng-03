FROM postgres:15-alpine

ENV POSTGRES_DB=orders_db
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password

COPY ./migrations /docker-entrypoint-initdb.d/

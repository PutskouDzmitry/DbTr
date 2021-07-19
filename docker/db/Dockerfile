FROM postgres:latest

COPY ./pkg/db/script/init.sql /docker-entrypoint-initdb.d/init.sql

EXPOSE 5432
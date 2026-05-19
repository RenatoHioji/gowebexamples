FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=admin
ENV MYSQL_USER=admin
ENV MYSQL_DATABASE=test

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 3306
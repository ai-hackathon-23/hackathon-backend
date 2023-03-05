FROM mysql:8.0

COPY mysql/init.sql /docker-entrypoint-initdb.d/

EXPOSE 3306

ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=mydb
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword

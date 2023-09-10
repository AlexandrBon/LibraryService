FROM mysql/mysql-server

COPY mysql/scripts/tables.sql /docker-entrypoint-initdb.d
COPY mysql/scripts/test_data.sql /docker-entrypoint-initdb.d

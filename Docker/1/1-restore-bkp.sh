#!/bin/bash
mkdir -p /tmp/wp
mkdir -p /tmp/mysql-db
#mkdir -p /tmp/docker-entrypoint-initdb.d
tar -C / -xvf ./all-databases.lzma.tar && docker compose -f docker-compose-recovery.yaml up -d
echo ""
echo ""
echo "Please, wait for MySQL-database recovery process complete and restart with regular status."
echo "(docker logs -f 1-db-1)"
echo "...and open URL: http://docker-host-ip:8080/"

#!/bin/bash
docker compose down --volumes
sudo rm -rf /tmp/wp
sudo rm -rf /tmp/mysql-db
sudo rm -rf /tmp/docker-entrypoint-initdb.d

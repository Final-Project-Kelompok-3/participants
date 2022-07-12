#!/usr/bin/env bash

docker container create --name mysql_db_docker --env MYSQL_ROOT_PASSWORD="root" --mount "type=volume,source=participants,destination=/var/lib/mysql" --network "host" -p 3306:3306 mysql:latest

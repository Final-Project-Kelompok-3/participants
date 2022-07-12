#!/usr/bin/env bash

docker build -t final-project-3/participants:1.0.0 .

docker container create --name db_mysql_container --env MYSQL_ROOT_PASSWORD="root" -p 45678:45678 --mount "type=bind,source=/home/participants/Dev/mount,destination=/var/lib/mysql" mysql:latest

docker run -d -it -p 8080:8080 --name=users --network="host"  final-project-3/participants:1.0.0

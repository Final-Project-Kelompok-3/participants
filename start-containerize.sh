#!/usr/bin/env bash

docker build -t final-project-3/participants:1.0.0 .

# docker container create --name mysql_db_docker --env MYSQL_ROOT_PASSWORD="root" --mount "type=volume,source=participants,destination=/var/lib/mysql" --network "host" -p 3306:3306 mysql:latest

docker run -d -it -p 8030:8030 --name=users --network="host"  final-project-3/participants:1.0.0

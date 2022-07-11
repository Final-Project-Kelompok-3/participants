#!/usr/bin/bash

docker build -t final-project-3/participants:1.0.0 .

docker run -d -it -p 45678:45678 --env MYSQL_ROOT_PASSWORD="root" --name=mysql_docker --add-host host.docker.internal:host-gateway mysql
docker run -d -it -p 8080:8080 --name=users  final-project-3/participants:1.0.0

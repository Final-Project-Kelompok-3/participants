#!/usr/bin/bash

docker build -t final-project-3/participants:1.0.0 .


docker run -d -it -p 8080:8080 --name=users  final-project-3/participants:1.0.0

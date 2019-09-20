#!/bin/bash

eval $(minikube docker-env)
./go_build.sh
docker build . -t sample-server
docker run -p 4000:80 -i sample-server

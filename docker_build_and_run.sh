#!/bin/bash

docker build . -t sample-server
docker run -i sample-server

#!/bin/bash

cd $(go env GOROOT)
cd src
GOOS=linux GOARCH=amd64 ./make.bash

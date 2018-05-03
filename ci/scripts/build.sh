#!/bin/bash -eux

cwd=$(pwd)

export GOPATH=$cwd/go

pushd $GOPATH/src/github.com/ONSdigital/dp-developer-site
  go run main.go
popd

cp -r $GOPATH/src/github.com/ONSdigital/dp-developer-site/assets/* build/
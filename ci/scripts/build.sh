#!/bin/bash -eux

cwd=$(pwd)

export GOPATH=$cwd/go

pushd $GOPATH/src/github.com/ONSdigital/dp-developer-site
  make build && cp build/dp-developer-site $cwd/build
popd

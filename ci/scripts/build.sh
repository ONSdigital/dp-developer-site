#!/bin/bash -eux

cwd=$(pwd)

export GOPATH=$cwd/go

pushd $cwd/dp-developer-site
  make build
popd

cp -r $cwd/dp-developer-site/assets/* build/
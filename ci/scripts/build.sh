#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-developer-site
  make build
popd

cp -r $cwd/dp-developer-site/assets/* build/
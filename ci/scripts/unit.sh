#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-developer-site
  make test
popd

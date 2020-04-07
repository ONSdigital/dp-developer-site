#!/bin/bash -eux

export cwd=$(pwd)

pushd $cwd/dp-developer-site
  make audit
popd
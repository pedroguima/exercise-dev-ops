#!/bin/bash

set -uex

export GOARCH=amd64
export GOOS=linux
export PKG=$1
export WD=$(pwd)
export binaryfiledir="$WD/ci/docker/$PKG/"
export binaryfile="${binaryfiledir}/${PKG}_${VERSION}_linux64"

cd src/$PKG
go get
go build -o $binaryfile .
strip $binaryfile 

if [ $PKG == "server" ]; then
  mkdir $binaryfiledir/testdata
  cp -a ../testdata/*.json $binaryfiledir/testdata
fi


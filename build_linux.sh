#!/bin/sh

cd src &&
echo Running tests:
go test ./... &&
if [ ! -d "../bin" ]; then
  echo Create output directory \'bin\'
  mkdir ../bin
fi &&
echo Building application
go build -ldflags="-s -w" -o ../bin/clnr &&
echo Done

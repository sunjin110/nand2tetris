#!/bin/sh

./compiler ./src/

rm -rf bin
mkdir bin
mv ./src/*.vm ./bin/

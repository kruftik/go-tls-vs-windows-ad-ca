#!/bin/bash

[ ! -d bin ] && mkdir bin

for C in client server http-server; do 
  go build -o bin/${C} ${C}.go
done

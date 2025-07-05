#!/bin/bash

source ~/.bashrc
cd app
go get
GIN_MODE=release go build -o ./build/appstarter
./build/appstarter
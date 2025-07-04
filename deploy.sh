#!/bin/bash

source ~/.bashrc
cd app
GIN_MODE=release go build -o ./build/appstarter
./build/appstarter
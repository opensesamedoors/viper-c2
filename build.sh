#!/bin/bash
mkdir -p dist

# build server
cd server 
go build .
GOOS=windows go build .
cd ..
cp server/server dist/server 
cp server/server.exe dist/server.exe
echo "Built the server"

# pack libary
zip dist/lib.zip lib/*.py lib/viper/*.py
echo "Packed the libary"

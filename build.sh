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

# pack python libary
zip dist/viperpy.zip lib/viperpy/*.py lib/viperpy/viper/*.py
echo "Packed the python libary"

#!/bin/sh

# generate jwt key
openssl genrsa -out jwtRS2048.key 2048
openssl rsa -in jwtRS2048.key -outform PEM -pubout > jwtRS2048.key.pub
mv jwtRS2048.key* build/

# generate protocol buffer
protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      protobuf/gogoatalk.proto

# copy account
cp ./server/accounts.json ./build/

# build server
env GOOS=linux GOARCH=amd64 go build -o ./build/server_linux ./server/server.go
env GOOS=darwin GOARCH=amd64 go build -o ./build/server_macos ./server/server.go
env GOOS=windows GOARCH=amd64 go build -o ./build/server_win.exe ./server/server.go

# build client
env GOOS=linux GOARCH=amd64 go build -o ./build/client_linux ./client/client.go
env GOOS=darwin GOARCH=amd64 go build -o ./build/client_macos ./client/client.go
env GOOS=windows GOARCH=amd64 go build -o ./build/client_win.exe ./client/client.go

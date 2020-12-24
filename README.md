# Gogoatalk - Simple TUI chat written in Go
This is an assignment for Network class.
The client may not work in windows environment.
This can be useful as a tutorial of go-grpc.

## Feature
- TUI based chat
- cross platform(windows, mac, linux)
- g-rpc
- jwt
- session

## Prerequisite
* protoc `brew install protobuf`

## How to Use
```sh
export GO111MODULE=on  # Enable module mode
go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
./build.sh
./server
# open new terminal
./client -h [server ip]
```

# Static account
You can add more users in server/accounts.json
```json
{
  "accounts": [
    {
      "userid": "ZVU30I96EG",
      "username": "alice",
      "password": "alice"
    }, {
      "userid": "CMZ95X73E7",
      "username": "jenny",
      "password": "jenny"
    }, {
      "userid": "G2H8CIJ24I",
      "username": "bob",
      "password": "bob"
    }, {
      "userid": "CQEO6BAOXO",
      "username": "tom",
      "password": "tom"
    }, {
      "userid": "8LPDBMHMIZ",
      "username": "jack",
      "password": "jack"
    }, {
      "userid": "JAKMPGYN73",
      "username": "kevin",
      "password": "kevin"
    }, {
      "userid": "GC7W2FEMLH",
      "username": "robin",
      "password": "robin"
    }, {
      "userid": "SHI42SLML5",
      "username": "james",
      "password": "james"
    }, {
      "userid": "20JA1AXM17",
      "username": "scott",
      "password": "scott"
    }, {
      "userid": "H68ZFXD4WT",
      "username": "fred",
      "password": "fred"
    }, {
      "userid": "KRZ18EDDOZ",
      "username": "larry",
      "password": "larry"
    }, {
      "userid": "PM87PSTO3K",
      "username": "jeff",
      "password": "jeff"
    }, {
      "userid": "S3Y3VVAGM9",
      "username": "steve",
      "password": "steve"
    }, {
      "userid": "G31fZXF1AE",
      "username": "john",
      "password": "john"
    }, {
      "userid": "ZT125VQVB3",
      "username": "sam",
      "password": "sam"
    }
  ]
}
```


## Screenshot




DID Universal Resolver Driver for did:plc
=========================================

This is an extremely simple golang proxy for resolving DIDs from <https://plc.directory>.


## Develop

```sh
go run main.go
```

then resolve a DID:

```sh
curl -X GET http://localhost:8000/1.0/identifiers/did:plc:yk4dd2qkboz2yv6tpubpc6co
curl -X GET http://localhost:8000/1.0/identifiers/did:plc:44ybard66vv44zksje25o7dz
```

## Docker

```sh
docker build -t bnewbold/uni-resolver-driver-did-plc -f Dockerfile .
docker push bnewbold/uni-resolver-driver-did-plc:latest
```

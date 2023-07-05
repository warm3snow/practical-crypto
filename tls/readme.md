## TLS

- tls1.2
- tls1.3
- gmtls1.1 TODO
- tlsclient

## Server

```golang
tls1.2/tls1.2.go
tls1.3/tls1.3.go
gmtls1.1/gmtls1.1.go //TODO

go run xxx.go //to start the server with tls of specified version
```

*Note: the default server key and cert is in testdata dir relative to the project.*

## Client
`tlsclient.go` is a https `client`, support: 
- tls1.0
- tls1.2
- tls1.3
- gmtls1.1

Usage: 
```sh
Usage of ./tlsclient:
  -addr string
        addr (default "https://localhost:8443")
  -skipVerify
        skipVerify (default true)
  -tlsVersion string
        specified tls version (default "1.2")
```

## Example
Example1:
```sh
# start the server with tls1.2
go run tls1.2/tls1.2.go

# start the client with tls1.2
go run tlsclient.go -addr https://localhost:8443 -tlsVersion 1.2 -skipVerify true
```

Example2:
```sh
# start the server with tls1.3
go run tls1.3/tls1.3.go

# start the client with tls1.3
go run tlsclient.go -addr https://localhost:8443 -tlsVersion 1.3 -skipVerify true
```

Example3:
```sh
# TODO start the server with gmtls1.1
```

*Note: tls server and client should run in different terminal*


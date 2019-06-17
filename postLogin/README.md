# PostLogin Service

This is the PostLogin service

Generated with

```
micro new RentHouseWeb/postLogin --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.postLogin
- Type: srv
- Alias: postLogin

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./postLogin-srv
```

Build a docker image
```
make docker
```
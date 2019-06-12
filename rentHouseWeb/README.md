# RentHouseWeb Service

This is the RentHouseWeb service

Generated with

```
micro new RentHouseWeb/rentHouseWeb --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.rentHouseWeb
- Type: web
- Alias: rentHouseWeb

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
./rentHouseWeb-web
```

Build a docker image
```
make docker
```
# GetSession Service

This is the GetSession service

Generated with

```
micro new RentHouseWeb/getSession --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.getSession
- Type: srv
- Alias: getSession

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
./getSession-srv
```

Build a docker image
```
make docker
```
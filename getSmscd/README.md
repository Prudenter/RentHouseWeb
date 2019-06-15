# GetSmscd Service

This is the GetSmscd service

Generated with

```
micro new RentHouseWeb/getSmscd --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.getSmscd
- Type: srv
- Alias: getSmscd

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
./getSmscd-srv
```

Build a docker image
```
make docker
```
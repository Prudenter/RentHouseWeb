# GetImageCd Service

This is the GetImageCd service

Generated with

```
micro new RentHouseWeb/getImageCd --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.getImageCd
- Type: srv
- Alias: getImageCd

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
./getImageCd-srv
```

Build a docker image
```
make docker
```
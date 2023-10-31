# tsj

Template for REST API server made with Golang.

Features:
- Graceful stop
- Close connections before stop
- Response template error/success/success with pagination

## Installation

To install `tsj` package, you need to install Go.

1. You first need [Go](https://golang.org/) installed then you can use the below Go command to install `tsj`.

```sh
go get -u github.com/studiobflat/tsj
```

2. Import it in your code:

```go
import "github.com/studiobflat/tsj"
```

## Quick start

### Starting HTTP server

The example how to use `tsj` with to start REST api in [example](./example/example) folder.




          
# GoCache - Redis Clone in Go

A in-memory cache server implementation inspired by Redis, written in Go and compatible with Redis clients, clis and...

## Features

- TCP server implementation
- In-memory key-value storage
- Support for TTL (Time To Live) on keys
- Concurrent connections handling
- Thread-safe operations
- Supports RESP

## Supported Commands

Currently supports the following Redis-like commands:

- `SET key value [ttl]` - Set a key with a value and optional TTL in seconds
- `GET key` - Get the value of a key

## Installation

```bash
git clone https://github.com/vesal-j/gocache.git
cd gocache
go build
```

## Usage

### Starting the Server

```bash
./gocache
```

The server will start listening on port 6969 by default.

### Connecting to the Server

You can use netcat (nc) to connect to the server:

```bash
nc localhost 6969
```

### Example Commands

```
SET mykey myvalue
OK

SET mykey myvalue 10  # Set with 10 seconds TTL
OK

GET mykey
myvalue

# After 10 seconds
GET mykey
(nil)
```

## Architecture

- `internal/core/app.go` - Main application logic and TCP server implementation
- `internal/store` - In-memory storage implementation with TTL support
- `internal/command` - Command handlers (GET, SET)
- `internal/core/router.go` - Command routing and processing

## Features

1. **Concurrent Connections**: Uses goroutines to handle multiple client connections
2. **TTL Support**: Automatic key expiration with background cleanup
3. **Thread-safe Operations**: Uses mutex for safe concurrent access
4. **Command Router**: Extensible command routing system
5. **Clean Architecture**: Modular design for easy extension

## Requirements

- Go 1.23.1 or higher

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

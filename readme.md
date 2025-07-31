


          
# GoCache - Redis Clone in Go

A in-memory cache server implementation inspired by Redis, written in Go and compatible with Redis clients, clis and etc, this project built just for fun 😊

## Features

- TCP server implementation
- In-memory key-value storage
- Support for TTL (Time To Live) on keys
- Concurrent connections handling
- Thread-safe operations
- Supports RESP (Redis Serialization Protocol)

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

The server will start listening on port 6380 by default.

### Connecting to the Server

Since GoCache uses RESP (Redis Serialization Protocol), you should use `redis-cli` to connect to the server:

```bash
redis-cli -p 6380
```

Or connect to a specific host:

```bash
redis-cli -h localhost -p 6380
```

### Example Commands

Once connected with redis-cli, you can use standard Redis commands:

```bash
# Set a key-value pair
SET mykey myvalue
OK

# Set a key with TTL (expires in 10 seconds)
SET mykey myvalue EX 10
OK

# Get a value
GET mykey
myvalue

# Check if key exists
EXISTS mykey
(integer) 1

# Get TTL for a key
TTL mykey
(integer) 8

# After 10 seconds, the key expires
GET mykey
(nil)

# Test connection
PING
PONG

# Get server info
INFO
# ... server information ...

# Get database size
DBSIZE
(integer) 0
```

### Using with Other Redis Clients

Since GoCache implements the RESP protocol, it's compatible with most Redis clients:

- **Python**: `redis-py`
- **Node.js**: `redis` or `ioredis`
- **Java**: `Jedis`
- **C#**: `StackExchange.Redis`

Example with Python:

```python
import redis

r = redis.Redis(host='localhost', port=6380)
r.set('mykey', 'myvalue', ex=10)
value = r.get('mykey')
print(value)  # b'myvalue'
```

## Architecture

- `internal/core/app.go` - Main application logic and TCP server implementation
- `internal/store` - In-memory storage implementation with TTL support
- `internal/command` - Command handlers (GET, SET, etc.)
- `internal/core/router.go` - Command routing and processing
- `internal/utils` - RESP protocol utilities

## Features

1. **Concurrent Connections**: Uses goroutines to handle multiple client connections
2. **TTL Support**: Automatic key expiration with background cleanup
3. **Thread-safe Operations**: Uses mutex for safe concurrent access
4. **Command Router**: Extensible command routing system
5. **Clean Architecture**: Modular design for easy extension
6. **RESP Protocol**: Full Redis Serialization Protocol support for client compatibility

## Requirements

- Go 1.23.1 or higher
- redis-cli (for testing) or any Redis client library

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

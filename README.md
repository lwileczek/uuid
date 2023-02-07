# UUID
Yet another implimentation of UUIDs in Go.
A UUID is encoded as a 128-bit object. The full offical Description from 2005: [RFC 4122]


## Goal
To create a lambda function which returns UUIDs.
Eventually, I'd love to make that lambda free and open to everyone with some rate limits.

## UUIDs
### v1
The first version takes the node the code is running on plus time to make a "random" bit sequence
### v4
Use crypto/rand to create pseudo-random bytes and then set the version and variation bytes
### pseudo 
Use crypto/rand to create 128 pseudo-random bits and then format them as a UUID hex string

## Build
Like any go project, ensure you have Go downloaded (1.19+) and run `go build .`

# Fun Stuff
## Benchmarks
Benchmarks don't really mean anything here but it's interesting to note that since UUIDv1 is incrementing
the time or clock sequence which is kept in a struct, so we don't need to make a new allocation with multiple
requests. So it's crazy fast even though it's the least random we could really make.

```
go test -benchmem -bench=. 
goos: darwin
goarch: amd64
pkg: github.com/lwileczek/uuid
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
BenchmarkPseudoUUID-8            1116513              1034 ns/op              16 B/op          1 allocs/op
BenchmarkV1-8                   16075844                72.30 ns/op            0 B/op          0 allocs/op
BenchmarkV4-8                    1000000              1025 ns/op              16 B/op          1 allocs/op
PASS
ok      github.com/lwileczek/uuid    4.384s
```

## Related Projects and Links
  - Google's Go UUID package: https://github.com/google/uuid
  - Another Popular Go UUID package v1-5: https://github.com/satori/go.uuid
  - Apparently there is a v6-8: https://github.com/uuid6/uuid6go-proto
  - Very concise and easy to read v4 https://github.com/bitactro/UUIDv4
  - https://www.cryptosys.net/pki/uuid-rfc4122.html
  - https://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language



[RFC 4122]: https://www.rfc-editor.org/rfc/rfc4122
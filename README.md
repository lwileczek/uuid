![build status](https://github.com/lwileczek/uuid/actions/workflows/basic.yml/badge.svg)
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
### Empty/Nil
The empty or Nil UUID is all zeros: `00000000-0000-0000-0000-000000000000`

## Build
Like any go project, ensure you have Go downloaded (1.19+) and run `go build ./...`
To build the CLI only `CGO_ENABLED=0 go build ./cmd/cli -o uuid`.

## Related Projects and Links
  - Google's Go UUID package: https://github.com/google/uuid
  - Another Popular Go UUID package v1-5: https://github.com/satori/go.uuid
  - Apparently there is a v6-8: https://github.com/uuid6/uuid6go-proto
  - Very concise and easy to read v4 https://github.com/bitactro/UUIDv4
  - https://www.cryptosys.net/pki/uuid-rfc4122.html
  - https://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language
  - Website to get v1 & v4 UUIDs: https://www.uuidgenerator.net/


[RFC 4122]: https://www.rfc-editor.org/rfc/rfc4122

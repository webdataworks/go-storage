/*
Package storage intends to provide a unified storage layer for Golang.

Goals

- Production ready: high test coverage, enterprise storage software adaptation, semantic versioning, well documented.

- High performance: more code generation, less runtime reflect.

- Vendor agnostic: more generic abstraction, less internal details.

Examples

The most common case to use a Storager service could be following:

1. Init a storager.

    store, err := fs.NewStorager(pairs.WithWorkDir("/tmp"))
	if err != nil {
		log.Fatalf("service init failed: %v", err)
	}

2. Use Storager API to maintain data.

	var buf bytes.Buffer

	n, err := store.Read("path/to/file", &buf)
	if err != nil {
		log.Printf("storager read: %v", err)
	}

*/
package storage

// We used to insert "-tags tools" here, but go-bindata doesn't support the new build
// tag that introduced in go 1.17. So we remove the tags here.
// In the further, we will move to go 1.16 embed to solve this problem.
//go:generate go run github.com/kevinburke/go-bindata/go-bindata -nometadata -o ./cmd/definitions/bindata/bindata.go -pkg bindata ./definitions
//go:generate go run -tags tools ./cmd/definitions

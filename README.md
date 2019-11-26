
This works on GNU/Linux only, as the link flags in `bindings.go` need to be modified for other platforms.  To get those flags, try this:

```
RUSTFLAGS='--print native-static-libs' cargo build --release --all
```

Installation

```
# Install rust and cbindgen and go

cargo build --release --all
cd pedersen-hash-ffi
./cbindgen
cd ..

# Verify things work
$ go build bindings.go

# Run sime test
$ go test bindings_test.go -v
=== RUN   TestPedersenHash
input:  [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
pedersen_hash_ffi called
output:  [216 21 67 136 189 3 97 66 78 166 184 57 243 109 73 57 140 151 139 155 161 138 12 184 161 12 54 29 20 83 138 15]
--- PASS: TestPedersenHash (0.09s)
PASS
ok      command-line-arguments  0.090s
```

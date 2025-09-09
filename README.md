# WASI bindings for Go

🟪 Centralized [WASI Preview 2](https://github.com/WebAssembly/WASI/tree/v0.2.7/wasip2) import bindings for Go

Looking for WASI Preview 1? See [v0.1.0](https://github.com/jcbhmr/go-wasi/tree/v0.1.0) \
Looking for WASI Preview 0? See [v0.0.0](https://github.com/jcbhmr/go-wasi/tree/v0.0.0)

<p align=center>
    TODO: Add image here
</p>

✅ Works with the official Go toolchain \
🤏 Works with the TinyGo compiler

- **[cli](https://pkg.go.dev/github.com/jcbhmr/go-wasi/cli)**
- **[clocks](https://pkg.go.dev/github.com/jcbhmr/go-wasi/clocks)**
- **[filesystem](https://pkg.go.dev/github.com/jcbhmr/go-wasi/filesystem)**
- **[http](https://pkg.go.dev/github.com/jcbhmr/go-wasi/http)**
- **[io](https://pkg.go.dev/github.com/jcbhmr/go-wasi/io)**
- **[random](https://pkg.go.dev/github.com/jcbhmr/go-wasi/random)**
- **[sockets](https://pkg.go.dev/github.com/jcbhmr/go-wasi/sockets)**

## Development

Development dependencies not tracked by Go:

- https://crates.io/crates/wasm-tools
- https://wasmtime.dev/
- https://crates.io/crates/wkg
- https://tinygo.org/

The official Go toolchain uses 64-bit pointers for `GOARCH=wasm`. TinyGo uses 32-bit pointers for `wasip1` and `wasip2` targets. `//go:wasmimport` and `//go:wasmexport` top-level signature pointers are actually treated as 32-bit pointers on the official Go toolchain. 😕 This means that we effectively need a conversion from native 64-bit pointer-containing structs to `uint32`-containing structs for the official Go toolchain but we can skip that step for TinyGo. `wit-bindgen-go` does **not** account for this and just assumes 32-bit pointers and TinyGo usage.

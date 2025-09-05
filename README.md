# WASI bindings for Go

🟪 Centralized WASI Preview 2 import bindings for Go

## Installation

```sh
go get github.com/jcbhmr/go-wasi@v0.2.0
```

**Why not v0.2.7?** \
Because you probably want to `//go:wasmimport` the WASI v0.2.0 interfaces. Not all hosts will provide WASI v0.2.7 imports to your guest WASM component yet. WASI v0.2.0 is (relatively) widely used and supported.

<details><summary>Looking for WASI Preview 1? Use v0.1.0.</summary>

```sh
go get github.com/jcbhmr/go-wasi@v0.1.0
```

</details>

Looking for WASI Preview 0? Use v0.0.0.

```sh
go get github.com/jcbhmr/go-wasi@v0.0.0
```

⚠️ v0.0.0, v0.1.0, and v0.2.0 all share the **same Go module name**

## Usage

This Go module `//go:wasmimport`-s all the WASI Preview 2 interfaces defined in the [WebAssembly/WASI](https://github.com/WebAssembly/WASI) specification repository. The bindings are generated using [go.bytecodealliance.org/cmd/wit-bindgen-go](https://pkg.go.dev/go.bytecodealliance.org/cmd/wit-bindgen-go) with some postprocessing to make them more ergonomic to consume as a standalone library.

What this means is that this package is a drop-in replacement for your `wit-bindgen-go generate`-ed `internal/wasi/` folder. Add something like this to wherever your `wit-bindgen-go generate` step is:

```diff
 //go:generate go tool wit-bindgen-go generate --out ./internal/ ./wit/
+//go:generate go tool jet "[\w\.-]+/internal/wasi/" "github.com/jcbhmr/go-wasi/" ./internal/
+//go:generate rm -rf ./internal/wasi/
```

[github.com/NicoNex/jet](https://pkg.go.dev/github.com/NicoNex/jet) is a find-and-replace CLI tool written in Go which means its easy to integrate with `go tool`. You can use [`find` & `sed`](https://stackoverflow.com/a/905161) or [`sd`](https://github.com/chmln/sd) if you prefer.

**Why not just use my project's local generated WASI bindings at `github.com/octocat/awesomelib/internal/wasi/...?** \
Because then your types are specific to _your module_. Your function `UseInputStream(streams.InputStream)` will accept `github.com/octocat/awesomelib/internal/wasi/io/streams.InputStream` but not `github.com/awesomesauce1000/coolapp/internal/wasi/io/streams.InputStream`. This is bad if you expose the WASI types in any way as part of your module's public API. Having a shared `github.com/jcbhmr/go-wasi/io/streams.InputStream` means that both `awesomelib` and `coolapp` can use the same type and pass WASI types between each other safely.

Hopefully someday Go will get a `syscall/wasi` or `golang.org/x/sys/wasi` package and this project can be deprecated. 🙏

You are encouraged to use this package in conjunction with a Go runtime binds to WASI Preview 2 like [TinyGo](https://tinygo.org/).

## Development

Development dependencies not tracked by Go:

- https://crates.io/crates/wasm-tools
- https://wasmtime.dev/
- https://crates.io/crates/wkg
- https://tinygo.org/

You might need to generate a local Go IDE configuration by running the VS Code command `>TinyGo target` and choosing `wasip2`. Don't commit the generated `settings.json`; it's got a machine-specific `GOROOT` path in it.

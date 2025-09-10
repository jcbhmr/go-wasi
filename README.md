# WASI bindings for Go

🧪 Centralized WASI Preview 0 import bindings for Go

_You might be looking for the [WASI Preview 1 v0.1.0](https://github.com/jcbhmr/go-wasi/tree/v0.1.0) version of this module_

<p align=center>
    TODO: Add image here
</p>

## Installation

```sh
go get github.com/jcbhmr/go-wasi@v0.0.0
```

⚠️ This module only works with WebAssembly targets.

## Usage

**When should I be using this module?** When you are compiling Go code for WASM and you want to **explicitly use WASI Preview 0 syscall-like functions** in your WASI application.

```go
package main

import (
    "fmt"

    "github.com/jcbhmr/go-wasi"
)

func main() {
    sizes, err := wasi.ArgsSizesGet()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("argc: %d, argvBufSize: %d\n", sizes.F0, sizes.F1)
}
```

## Development

Development dependencies not tracked by Go:

- https://crates.io/crates/wasm-tools
- https://wasmtime.dev/
- https://tinygo.org/

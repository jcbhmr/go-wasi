package wasi_test

import (
	"os"
	"os/exec"
	"testing"
)

func runCode(t *testing.T, code string) string {
	t.Helper()
	err := os.MkdirAll(".out", 0o755)
	if err != nil {
		t.Fatalf("could not make directory %q recursively: %v", ".out", err)
	}
	err = os.WriteFile(".out/code.go", []byte(code), 0o666)
	if err != nil {
		t.Fatalf("failed to write %q to %q: %v", code, ".out/code.go", err)
	}
	cmd := exec.CommandContext(t.Context(), "go", "build", "-o", ".out/code.wasm", ".out/code.go")
	cmd.Env = append(os.Environ(), "GOOS=wasip1", "GOARCH=wasm")
	t.Logf("Running %q", cmd)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		t.Logf("%s", output)
	}
	if err != nil {
		t.Fatalf("failed to execute %q: %v", cmd, err)
	}
	cmd = exec.CommandContext(t.Context(), "wasmtime", ".out/code.wasm")
	t.Logf("Running %q", cmd)
	output, err = cmd.CombinedOutput()
	if len(output) > 0 {
		t.Logf("%s", output)
	}
	if err != nil {
		t.Fatalf("failed to execute %q: %v", cmd, err)
	}
	return string(output)
}

func TestBase(t *testing.T) {
	runCode(t, `package main

import (
	"fmt"
	"unsafe"

	_ "github.com/jcbhmr/go-wasi"
)

func main() {
	fmt.Println("Hello!")
	fmt.Printf("Size of uintptr: %d bytes or %d bits", unsafe.Sizeof(uintptr(0)), unsafe.Sizeof(uintptr(0)) * 8)
}
`)
}

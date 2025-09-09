package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var outFlag = flag.String("out", "", "output directory")
var witFlag = flag.String("wit", "./wit/", "wit directory")
var cmFlag = flag.String("cm", "", "cm import path")
var hoistFlag = flag.String("hoist", "", "which path to hoist out of generated bindings")

func main() {
	flag.Parse()
	if *hoistFlag == "" {
		log.Fatal("-hoist <path> flag must be present")
	}

	log.Printf("Removing %q recursively", "./.out/wit-bindgen-go-wrapper/")
	err := os.RemoveAll("./.out/wit-bindgen-go-wrapper/")
	if err != nil {
		log.Fatalf("failed to remove %q recursively: %v", "./.out/wit-bindgen-go-wrapper/", err)
	}

	args := []string{"tool", "wit-bindgen-go", "generate", "--out", "./.out/wit-bindgen-go-wrapper/"}
	if *cmFlag != "" {
		args = append(args, "--cm", *cmFlag)
	}
	args = append(args, *witFlag)
	cmd := exec.Command("go", args...)
	log.Printf("Running %q", cmd)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		log.Printf("%s", output)
	}
	if err != nil {
		log.Fatalf("failed to run %q: %v", cmd, err)
	}

	hoistFrom := filepath.Join("./.out/wit-bindgen-go-wrapper/", *hoistFlag)
	hoistFromImportPathOld := "/" + path.Clean(path.Join(".out/wit-bindgen-go-wrapper", *hoistFlag)) + "/"
	entries, err := os.ReadDir(hoistFrom)
	if err != nil {
		log.Fatalf("failed to read directory %q: %v", hoistFrom, err)
	}
	for _, e := range entries {
		eRelativeName := filepath.Join(hoistFrom, e.Name())

		log.Printf("Removing %q recursively", e.Name())
		err := os.RemoveAll(e.Name())
		if err != nil {
			log.Fatalf("failed to remove %q recursively", err)
		}

		log.Printf("Copying %q to %q", eRelativeName, e.Name())
		err = os.CopyFS(e.Name(), os.DirFS(eRelativeName))
		if err != nil {
			log.Fatalf("failed to copy %q to %q", eRelativeName, e.Name())
		}

		err = fs.WalkDir(os.DirFS("./"), e.Name(), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			isGo, err := filepath.Match("*.go", d.Name())
			if err != nil {
				panic(fmt.Errorf("failed to match %q against %q: %w", d.Name(), "*.go", err))
			}
			if !isGo {
				return nil
			}

			code, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %q: %w", path, err)
			}
			code = bytes.Replace(code, []byte("DO NOT EDIT.\n"), []byte("DO NOT EDIT.\n\n//go:build wasm\n"), 1)
			code = bytes.ReplaceAll(code, []byte(hoistFromImportPathOld), nil)
			err = os.WriteFile(path, code, 0o666)
			if err != nil {
				return fmt.Errorf("failed to write %d bytes to file %q: %w", len(code), path, err)
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Removing %q", filepath.Join(e.Name(), "empty.s"))
		err = os.Remove(filepath.Join(e.Name(), "empty.s"))
		if err != nil {
			log.Fatalf("failed to remove %q: %v", filepath.Join(e.Name(), "empty.s"), err)
		}
	}
}

## Repository overview

This repository is a collection of small Go example modules and exercises (tutorial-style). Most folders contain standalone examples (often `package main`) demonstrating language features. The root module is `example.local/demo` (see `go.mod`), and several subfolders have their own `go.mod` (treat them as independent modules).

Key folders/files to reference
- `HelloWorld.go` — combined examples of functions, constants, and `main` usage.
- `DEMO/main.go` + `DEMO/funcs.go` — example of separating `main` and helper functions (useful pattern to follow).
- `UnitTesting/` — a library module (`module calctest2`) with `calc.go` and `calc_test.go`; shows how tests import the local module.
- `Testing/` — simple `package main` test example (`package_test.go`).
- `Channels/chan.go`, `Arrays/arr.go`, `Functions/func.go` — focused examples demonstrating idiomatic snippets used across the repo.

Big-picture architecture
- Not a single monolith — it's a learning repo comprised of many small, mostly independent modules. Expect many `package main` examples that are runnable independently.
- Modules: several directories include their own `go.mod`. Treat each as a self-contained module when building/testing.

Developer workflows (concrete commands)
- Run a single example (from project root):
  - cd into the example folder and run with PowerShell:
    ```powershell
    cd .\DEMO; go run .
    ```
  - Or run a specific file:
    ```powershell
    go run .\HelloWorld.go
    ```
- Build an example (produce an executable):
  ```powershell
  cd .\DEMO; go build -o demo.exe .
  .\demo.exe
  ```
- Run tests for a module (recommended):
  ```powershell
  cd .\UnitTesting; go test ./...
  ```
- Run tests for the whole repo (note: due to multiple go.mod files, this may require per-module runs):
  ```powershell
  go test ./...    # works when modules are arranged under a single module, otherwise run per-folder
  ```

Project-specific conventions and gotchas
- Multiple module roots: many subfolders include `go.mod`. When running `go` commands, prefer `cd` into the folder and run the command there. `go test ./...` from the root may not traverse into nested modules as you expect.
- Module names are sometimes used directly in tests/imports: `UnitTesting/go.mod` declares `module calctest2` and tests import `calctest2` (see `UnitTesting/calc_test.go`). Keep the module path consistent when moving files.
- Examples are intentionally simple and educational; some functions (e.g., `UnitTesting/calc.go`) have odd behaviors (it prints multiplication result while returning a sum). When refactoring, keep or clarify such side-effects.

How code is organized
- Most example files are `package main` and contain a `main()` function. These are runnable demos, not libraries.
- Library-style code (to import in tests) appears under `UnitTesting/` (`package calc`) — use that folder when adding reusable functions.

Integration points & external dependencies
- This workspace uses only the Go standard library. No external modules are referenced in provided `go.mod` files.

Examples to copy/paste
- Run the DEMO example:
  ```powershell
  cd .\DEMO; go run .
  ```
- Run UnitTesting tests (verifies `calc` behavior):
  ```powershell
  cd .\UnitTesting; go test
  ```

If you modify tests or module names
- Keep `module` in `go.mod` and import paths aligned. If you rename a module, update imports in tests that refer to the module name directly (e.g., `calc_test` imports `calctest2`).

What an AI agent should do first
1. Recognize each directory as an independent example module unless it lacks `go.mod`.
2. Run one example (`DEMO`) and one test module (`UnitTesting`) to confirm the local environment and Go toolchain behavior.
3. When adding new code, prefer adding small, focused files under a new subfolder with its own `go.mod` to keep the repo consistent.

Where to look for more context
- Root `go.mod` (Go version: 1.25.1) and each subfolder's `go.mod`.
- `DEMO/funcs.go` demonstrates splitting helper functions from `main`.

If something is unclear
- Ask which example or module you should run first. If tests fail, include `go test` output and which folder you ran it from.

---
This file was generated to help AI coding agents quickly become productive in this learning-focused Go workspace. Ask for clarifications or to expand examples for any specific folder.

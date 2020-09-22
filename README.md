# Example GitHub Actions workflow for Go ![CI](https://github.com/amwolff/example-github-actions-go/workflows/CI/badge.svg)

Inspired by the awesome [GitHub Actions for Go](https://github.com/mvdan/github-actions-golang).

## The workflow

```yaml
$ cat .github/workflows/go.yml
name: CI

on:
  push:
    branches: [ main ]
    tags: [ v* ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    name: Build & Test
    runs-on: ${{ matrix.os }}
    steps:
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: ${{ matrix.go-version }}

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Cache ~/go/pkg/mod
        uses: actions/cache@v2
        with:
          path: ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ runner.os }}-go-

      - name: Build
        run: go build -a -v ./...

      - name: Test
        run: go test -a -race -v ./...
    strategy:
      matrix:
        go-version: [ 1.15.x ]
        os: [ ubuntu-latest ]
  lint:
    name: Lint
    runs-on: ubuntu-latest
    steps:
      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: golangci-lint
        uses: golangci/golangci-lint-action@v2
        with:
          version: v1.31
          args: -E goimports
```

## Summary

This workflow

-[x] uses new GitHub's default branch name;
-[x] uses caching;
-[x] uses [golangci-lint](https://golangci-lint.run/) with default settings and `goimports` enabled;
-[x] is easy to use and extend!

You can add more Go versions and operating systems into the [strategy matrix](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategymatrix), modify Go compiler flags and play around with `golangci-lint` linters.

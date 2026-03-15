# Go CLI Text Utilities

A Go utility project that provides two core text-processing functions: word-frequency counting and palindrome detection, with test coverage for both modules.

## Phase Position

- **A2SV Go Phase:** Task 2 (Foundation+Testing)
- **Previous Project:** `go-cli-student-grade-calculator`
- **Next Project:** `go-cli-library-management-system`

## Features

- Case-insensitive word frequency analysis
- Punctuation-tolerant palindrome checker
- Focused unit tests for both features

## Tech Stack

- Go
- Go testing package (`go test`)

## Project Structure

```text
.
├── main.go
├── palindrome_test.go
├── wordfreq_test.go
├── go.mod
└── Readme.md
```

## Run

```bash
go run main.go
```

## Test

```bash
go test -v ./...
```

## Learning Outcomes

- String normalization and tokenization
- Utility-style function decomposition
- Early-stage testing discipline in Go

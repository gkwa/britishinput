# BritishInput

A Go CLI tool for testing GitHub Actions pipeline resilience and retry logic. BritishInput helps validate pipeline configurations by providing deterministic and random exit status behaviors.

## Purpose

BritishInput was created to help test GitHub Actions pipeline retry mechanisms and failure handling. When running long pipelines, transient failures often occur that can be resolved by retrying. This tool helps validate that your pipeline configuration correctly handles such scenarios by:

- Providing predictable success/failure exit statuses
- Simulating random failures with 50/50 probability
- Enabling testing of retry logic and failure cascades

## Installation

```bash
# Clone the repository
git clone https://github.com/gkwa/britishinput.git

# Build the binary
cd britishinput
go build
```

## Usage

BritishInput provides three modes of operation:

1. Default mode (50/50 chance of success/failure):

```bash
./britishinput
```

2. Always fail:

```bash
./britishinput fail
```

3. Always succeed:

```bash
./britishinput pass
```

Example usage with exit status checking:

```bash
$ ./britishinput; echo $?
0  # or 1 (random)

$ ./britishinput fail; echo $?
1  # always fails

$ ./britishinput pass; echo $?
0  # always succeeds
```

## GitHub Actions Integration

BritishInput is particularly useful for testing GitHub Actions pipeline configurations. Here's an example workflow that demonstrates retry logic:

```yaml
jobs:
  test-retry:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: ">=1.21"
      - name: Build
        run: go build -v
      - name: Run with potential failure
        run: ./britishinput
```

The tool can help validate:

- Retry mechanisms
- Failure cascades
- Pipeline recovery logic
- Complex conditional job execution

## Features

- Deterministic success/failure modes
- Random 50/50 success probability
- Simple CLI interface
- Configurable via environment variables and config file
- JSON logging support
- Verbose mode for debugging

## Version Information

View the current version:

```bash
./britishinput version
```

# Scale17x Talk - Two Binaries are Better than one?

`Building Golang and Python Static Binaries`

This repo is the boilerplate for building a Golang Binary with an embedded static python binary.

## Requirements

### Darwin Builds

* GO 1.11 (`go mod`)
* make
* python3
* [pyinstaller](https://pyinstaller.readthedocs.io/en/v3.3.1/operating-mode.html)
* [fileb0x](https://github.com/UnnoTed/fileb0x)
  
### Linux Builds

* make
* docker

## Usage

```bash
Makefile Help

make [targets...]
Targets:
  build-linux:            Build Both the Golang and Python Binaries (Platform: Linux)
  build-linux-python:     Build the Python Binary (Platform: Linux)
  build-linux-golang:     Build the Golang Binary (Platform: Linux)
  runtime-linux:          Docker Container to run Built Linux Binaries
  build-darwin:           Build Both the Golang and Python Binaries (Platform: Darwin)
  build-darwin-python:    Build the Python Binary (Platform: Darwin)
  build-darwin-golang:    Build the Golang Binary (Platform: Darwin)
  clean:                  Clean the bin directory
```
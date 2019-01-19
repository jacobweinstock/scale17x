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
  build-linux:            Build the Both the Golang and Python Binaries (Platform: Linux)
  build-python-linux:     Build the Python Binary (Platform: Linux)
  build-golang-linux:     Build the Golang Binary (Platform: Linux)
  runtime-linux:          Docker Container to run Built Linux Binaries
  build-darwin:           Build the Both the Golang and Python Binaries (Platform: Darwin)
  build-python-darwin:    Build the Python Binary (Platform: Darwin)
  build-golang-darwin:    Build the Golang Binary (Platform: Darwin)
  clean:                  Clean bin directory
```
GOLANG_BINARY := scale17x
PYTHON_BINARY := scale17x-py
.DEFAULT_GOAL := help


.PHONY: build-linux
build-linux: build-linux-python build-linux-golang	## Build Both the Golang and Python Binaries (Platform: Linux)

.PHONY: build-linux-python
build-linux-python:	## Build the Python Binary (Platform: Linux)
		mkdir -p bin/linux/python
		(cd python && docker build --compress --force-rm -t local/$(GOLANG_BINARY):python-linux -f Dockerfile .)
		docker run --rm -v $(PWD)/bin/linux/python:/tmp/bin local/$(GOLANG_BINARY):python-linux sh -c 'cp -a /home/bin/linux/python/$(PYTHON_BINARY) /tmp/bin/$(PYTHON_BINARY)'
		cp -a bin/linux/python/$(PYTHON_BINARY) golang/extmodules/$(PYTHON_BINARY)

.PHONY: build-linux-golang
build-linux-golang:	## Build the Golang Binary (Platform: Linux)
		mkdir -p bin/linux/golang
		(cd golang && docker build --compress --force-rm -t local/$(GOLANG_BINARY):linux-golang -f Dockerfile .)
		docker run --rm -v $(PWD)/bin/linux/golang:/tmp/bin local/$(GOLANG_BINARY):linux-golang sh -c 'cp -a /home/$(GOLANG_BINARY) /tmp/bin/$(GOLANG_BINARY)'

.PHONY: runtime-linux
runtime-linux:		## Docker Container to run Built Linux Binaries
		docker run -it --rm -w /app -v $(PWD)/bin/linux/golang/:/app ubuntu:18.04 bash

.PHONY: build-darwin
build-darwin: build-darwin-python build-darwin-golang	## Build Both the Golang and Python Binaries (Platform: Darwin)

.PHONY: build-darwin-python
build-darwin-python:	## Build the Python Binary (Platform: Darwin)
		mkdir -p bin/darwin/python
		pyinstaller ./python/scale17x/$(PYTHON_BINARY).py --onefile --clean --distpath bin/darwin/python/
		rm -rf build $(PYTHON_BINARY).spec

.PHONY: build-darwin-golang
build-darwin-golang:	## Build the Golang Binary (Platform: Darwin)
		mkdir -p bin/darwin/golang
		cp -a bin/darwin/python/$(PYTHON_BINARY) golang/extmodules/$(PYTHON_BINARY)
		(cd golang && fileb0x extmodules.yaml)
		(cd golang && GOOS=darwin GOARCH=amd64 go build -o ../bin/darwin/golang/$(GOLANG_BINARY)-darwin main.go)

.PHONY: clean
clean:	## Clean the bin directory
		@rm -rf bin/*

.PHONY: help
help:
		@echo
		@echo "Makefile Help\n\nmake [targets...]\nTargets:"
		@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/:.*##/: ##/' | sed -e 's/^/  /' | column -s## -t
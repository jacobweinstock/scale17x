GOLANG_BINARY := scale17x
PYTHON_BINARY := scale17x-py
.DEFAULT_GOAL := help


.PHONY: build-linux
build-linux: build-linux-python build-linux-golang	## Build Both the Golang and Python Binaries (Platform: Linux)

.PHONY: build-linux-python
build-linux-python:	## Build the Python Binary (Platform: Linux)
		mkdir -p bin/linux/python
		(cd python && docker build --compress --force-rm -t local/$(GOLANG_BINARY):linux-python -f Dockerfile .)
		docker rm -f cont1 || true && docker create --name cont1 local/${GOLANG_BINARY}:linux-python && docker cp cont1:/home/bin/linux/python/${PYTHON_BINARY} bin/linux/python/${PYTHON_BINARY} && docker rm -f cont1
		cp -a bin/linux/python/$(PYTHON_BINARY) golang/extmodules/$(PYTHON_BINARY)

.PHONY: build-linux-golang
build-linux-golang:	## Build the Golang Binary (Platform: Linux)
		mkdir -p bin/linux/golang
		(cd golang && docker build --compress --force-rm -t local/$(GOLANG_BINARY):linux-golang -f Dockerfile .)
		docker rm -f cont2 || true && docker create --name cont2 local/${GOLANG_BINARY}:linux-golang && docker cp cont2:/home/${GOLANG_BINARY} bin/linux/golang/${GOLANG_BINARY} && docker rm -f cont2

.PHONY: runtime-linux
runtime-linux:		## Docker Container to run Built Linux Binaries
		docker run -it --rm -p8080 -w /app -v $(PWD)/bin/linux/golang/:/app ubuntu:18.04 bash

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
		@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/:.*##/: ##/' | sed -e 's/^/  /' | column -s## -t | sort
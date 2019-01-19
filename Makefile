GOLANG_BINARY := scale17x
PYTHON_BINARY := scale17x-py

.PHONY: build-linux
build-linux: build-python-linux build-golang-linux

.PHONY: build-python
build-python-linux:
		mkdir -p bin/linux/python
		(cd python && docker build --compress --force-rm -t local/$(GOLANG_BINARY):python-linux -f Dockerfile .)
		docker run --rm -v $(PWD)/bin/linux/python:/tmp/bin local/$(GOLANG_BINARY):python-linux sh -c 'cp -a /home/bin/linux/python/$(PYTHON_BINARY) /tmp/bin/$(PYTHON_BINARY)'
		cp -a bin/linux/python/$(PYTHON_BINARY) golang/extmodules/$(PYTHON_BINARY)

.PHONY: build-golang-linux
build-golang-linux:
		mkdir -p bin/linux/golang
		(cd golang && docker build --compress --force-rm -t local/$(GOLANG_BINARY):linux-golang -f Dockerfile .)
		docker run --rm -v $(PWD)/bin/linux/golang:/tmp/bin local/$(GOLANG_BINARY):linux-golang sh -c 'cp -a /home/$(GOLANG_BINARY) /tmp/bin/$(GOLANG_BINARY)'

.PHONY: runtime-linux
runtime-linux:
		docker run -it --rm -w /app -v $(PWD)/bin/linux/golang/:/app ubuntu:18.04 bash

.PHONY: build-darwin
build-darwin: build-python-darwin build-golang-darwin

.PHONY: build-python-darwin
build-python-darwin:
		mkdir -p bin/darwin/python
		pyinstaller ./python/scale17x/$(PYTHON_BINARY).py --onefile --clean --distpath bin/darwin/python/
		rm -rf build $(PYTHON_BINARY).spec

.PHONY: build-golang-darwin
build-golang-darwin:
		mkdir -p bin/darwin/golang
		cp -a bin/darwin/python/$(PYTHON_BINARY) golang/extmodules/$(PYTHON_BINARY)
		(cd golang && fileb0x extmodules.yaml)
		(cd golang && GOOS=darwin GOARCH=amd64 /usr/local/Cellar/go/1.11.4/bin/go build -o ../bin/darwin/golang/$(GOLANG_BINARY)-darwin main.go)
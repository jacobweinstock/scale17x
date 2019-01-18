.PHONY: build-python
build-python-linux:
		docker build --compress --force-rm -t pyinstaller:local -f python/Dockerfile python/
		docker run -it --user root --rm -w /tmp -v $(PWD):/tmp pyinstaller:local bash -c "pyinstaller python/scale17x/run.py --paths=/tmp --onefile --clean --distpath bin/linux/; rm -rf build run.spec"

.PHONY: build-python-darwin
build-python-darwin:
		pyinstaller ./python/scale17x/run.py --onefile --clean --distpath bin/darwin/
		rm -rf build run.spec

FROM python:3.7

ARG PYTHON_BINARY=scale17x-py

WORKDIR /home/
ADD . /home/

RUN pip install pyinstaller PyCrypto>=2.4 && \
    pyinstaller scale17x/${PYTHON_BINARY}.py --paths=/home/python --onefile --clean --distpath /home/bin/linux/python/
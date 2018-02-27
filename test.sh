#!/bin/bash -e
FIFO=`mktemp -d`/fifo
mkfifo -m 666 ${FIFO}

go version
go run main.go ${FIFO}

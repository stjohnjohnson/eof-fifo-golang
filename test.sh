#!/bin/bash -e
FIFO=`mktemp -d`/fifo
mkfifo -m 666 ${FIFO}

go version
go run reader/main.go ${FIFO} &
go run writer/main.go ${FIFO} &

wait $(jobs -p)

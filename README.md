# EOF/FIFO bug with GoLang

I'm trying to understand why this code is working on Linux but not OSX.

## Usage

### Linux (docker)

```bash
$ docker build .
Sending build context to Docker daemon  43.52kB
Step 1/4 : FROM golang:1.10
 ---> 1c1309ff8e0d
Step 2/4 : WORKDIR /go/src/app
 ---> Using cache
 ---> d78b5cc2961d
Step 3/4 : COPY . .
 ---> f2d594442fb9
Step 4/4 : ENTRYPOINT /go/src/app/test.sh
 ---> Running in 1e69528c51b7
Removing intermediate container 1e69528c51b7
 ---> 2d23db5e8ff0
Successfully built 2d23db5e8ff0

$ docker run --rm 2d23db5e8ff0
go version go1.10 linux/amd64
STARTED /tmp/tmp.6if4ZSWnZ8/fifo
READER >> created
WRITER << opened: &{file:0xc42008c000}|<nil>
WRITER << encoder created
WRITER << written line0, <nil>
READER >> read 1 line: line0
WRITER << written line1, <nil>
READER >> read 1 line: line1
WRITER << written line2, <nil>
READER >> read 1 line: line2
WRITER << closed: <nil>
READER >> read finished: EOF
ALL DONE
```

### OSX

```bash
go version go1.10 darwin/amd64
STARTED /var/folders/k_/zdt3jb2j7gx1rnf379k_51gh0000gn/T/tmp.ksHGLSvs/fifo
READER >> created
WRITER << opened: &{file:0xc42009e000}|<nil>
WRITER << encoder created
WRITER << written line0, <nil>
READER >> read 1 line: line0
WRITER << written line1, <nil>
READER >> read 1 line: line1
WRITER << written line2, <nil>
READER >> read 1 line: line2
WRITER << closed: <nil>
^C
signal: interrupt
```

Note how the reader never gets the EOF.

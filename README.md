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
WRITER << opened /tmp/tmp.pghBSjMs2H/fifo: &{file:0xc4200840f0}|<nil>
WRITER << encoder created
READER >> created &{buf:[....] rd:0xc42008a020 r:0 w:0 err:<nil> lastByte:-1 lastRuneSize:-1}
WRITER << written line1, <nil>
READER >> read 1 line: line1
WRITER << written line2, <nil>
READER >> read 1 line: line2
WRITER << closed /tmp/tmp.pghBSjMs2H/fifo: <nil>
READER >> read finished: EOF
```

### OSX

```bash
$ ./test.sh
go version go1.10 darwin/amd64
WRITER << opened /var/folders/k_/zdt3jb2j7gx1rnf379k_51gh0000gn/T/tmp.4n5wTPm6/fifo: &{file:0xc4200840f0}|<nil>
WRITER << encoder created
READER >> created &{buf:[....] rd:0xc42000c030 r:0 w:0 err:<nil> lastByte:-1 lastRuneSize:-1}
WRITER << written line1, <nil>
READER >> read 1 line: line1
WRITER << written line2, <nil>
READER >> read 1 line: line2
WRITER << closed /var/folders/k_/zdt3jb2j7gx1rnf379k_51gh0000gn/T/tmp.4n5wTPm6/fifo: <nil>
^C%
signal: interrupt
```

Note how the reader never gets the EOF.

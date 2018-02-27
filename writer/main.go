package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fifo := os.Args[1]
	f, err := os.OpenFile(fifo, os.O_WRONLY, 0600)
	fmt.Printf("WRITER << opened %+v: %+v|%+v\n", fifo, f, err)
	if err != nil {
		panic(err)
	}

	fmt.Printf("WRITER << encoder created\n")

	time.Sleep(1 * time.Second)
	_, err = f.WriteString("line1\n")
	fmt.Printf("WRITER << written line1, %+v\n", err)

	time.Sleep(1 * time.Second)
	_, err = f.WriteString("line2\n")
	fmt.Printf("WRITER << written line2, %+v\n", err)

	time.Sleep(1 * time.Second)
	err = f.Close()
	fmt.Printf("WRITER << closed %+v: %+v\n", fifo, err)
}

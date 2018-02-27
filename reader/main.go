package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// @note This is used to get around bufio.Scanner erroring at large lines
func readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}

func main() {
	fifo := os.Args[1]

	t := time.AfterFunc(time.Minute, func() {
		fmt.Printf("No data in the first %s. Assuming catastophe.", time.Minute)
		os.Exit(1)
	})
	f, err := os.Open(fifo)
	t.Stop()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	fmt.Printf("READER >> created %+v\n", reader)

	line, readErr := readln(reader)

	for readErr == nil {
		fmt.Printf("READER >> read 1 line: %+v\n", line)
		line, readErr = readln(reader)
	}

	fmt.Printf("READER >> read finished: %+v\n", readErr)
}

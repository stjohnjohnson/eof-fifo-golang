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

func reader(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	fmt.Print("READER >> created\n")

	line, readErr := readln(reader)

	for readErr == nil {
		fmt.Printf("READER >> read 1 line: %+v\n", line)
		line, readErr = readln(reader)
	}

	fmt.Printf("READER >> read finished: %+v\n", readErr)
}

func writer(filePath string) {
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0600)

	fmt.Printf("WRITER << opened: %+v|%+v\n", f, err)
	if err != nil {
		panic(err)
	}

	fmt.Printf("WRITER << encoder created\n")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		_, err = f.WriteString(fmt.Sprint("line", i, "\n"))
		fmt.Printf("WRITER << written line%d, %+v\n", i, err)
	}

	time.Sleep(1 * time.Second)
	err = f.Close()
	fmt.Printf("WRITER << closed: %+v\n", err)
}

func main() {
	fifo := os.Args[1]

	fmt.Printf("STARTED %s\n", fifo)

	go writer(fifo)
	reader(fifo)

	fmt.Print("ALL DONE\n")
}

package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := os.Args[1]
	// fmt.Println(filename)
	// f, err := os.Open(string(filename))
	// b1 := make([]byte, 99)
	// n1, err := f.Read(b1)
	// check(err)
	// fmt.Println(n1, string(b1[:n1]))

	// dat, err := ioutil.ReadFile(string(filename))
	// check(err)
	// fmt.Print(string(dat))

	f, err := os.Open(string(filename))
	if err == nil {
		io.Copy(os.Stdout, f)
	} else {
		fmt.Println(err)
	}
}

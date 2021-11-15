package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "myapp/internal"
	"errors"
)

const (
	workPath = "/Users/u19502010/rebrainme/golandbase"
)

var (
	pathin         = workPath + "/module04/06_task/08_task_in.txt"
	result         string
	i              = 0
	limit          = 160
	errExcessLimit = "limit error"
)

func main() {
	readFile(pathin)
	fmt.Println("string count exceed limit, please read another file =) err: ", result)
}

func readFile(pathin string) {
	f, err := os.Open(pathin)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewReader(f)
	for {
		value, err := s.ReadString('\n')
		i++
		if err != nil {
			fmt.Println(blabla(value, i, err))
			break
		}
	}
	fmt.Printf("Total strings: %d\n", i)
}

func blabla(value string, i int, err error) error {
	switch {
	case i > limit:
		// result := internal.ErrParser2{
		// 	err,
		// 	limit: limit,
		// 	message: value}
		result = fmt.Sprintf("%s, limit: %d, last string: %s", errExcessLimit, limit, value)
	case err == io.EOF:
		result = "fuck EOF"
	default:
		result = "hz error"
	}
	return errors.New(result)
}
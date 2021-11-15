package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"myapp/internal"
	"errors"
)

const (
	workPath = "/Users/u19502010/rebrainme/golandbase"
)

var (
	pathin         = workPath + "/module04/06_task/08_task_in.txt"
	limit          = 150
	errExcessLimit = "limit error"
	errEOF         = errors.New("fuck EOF error")
)

func main() {
	err := readFile(pathin)
	if err != nil {
		if _, ok := err.(*internal.ErrParser); ok {
			fmt.Print("string count exceed limit, please read another file =) err: ", err.Error())
		} else if err == errEOF {
			fmt.Println(errEOF)
		} else {
			fmt.Println(err)
		}
	}
}

func readFile(pathin string) error {
    i := 0
	f, err := os.Open(pathin)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewReader(f)
	for {
		i++
		value, err := s.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return errEOF
			} else {
				return err
			}
		}
		if i > (limit-1){
			mylimit := internal.ErrParser{
				errExcessLimit,
				limit,
				value}
			return &mylimit
		}
	}
	return nil
}

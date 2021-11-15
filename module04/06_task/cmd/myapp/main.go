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
	result         string
	i              = 0
	limit          = 150
	errType        string
	errExcessLimit = "limit error"
	errEOF         = "fuck EOF error"
	errAnother     = "some another error"
)

func main() {
	readFile(pathin)
	switch errType{
	case errExcessLimit:
		fmt.Println("string count exceed limit, please read another file =) err: ", result)
	case errEOF:
		fmt.Println(errType)
	case errAnother:
		fmt.Println(errType)
	default:
		fmt.Println("No error")
	}
}

func readFile(pathin string) error {
	f, err := os.Open(pathin)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewReader(f)
	for {
		value, err := s.ReadString('\n')
		i++
		if i > limit{
			errType = errExcessLimit
			mylimit := internal.ErrParser{
				errType,
				limit,
				value}
            result = startParse(mylimit)
			return errors.New(result)
		}
		if err != nil {
			if err == io.EOF {
				errType = errEOF
				return errors.New(errType)
				break
			} else {
				errType = errAnother
				return errors.New(errType)
			}
		}
	}
	return nil
}

func startParse(mylimit internal.ErrParser) string {
	return mylimit.Error()
}

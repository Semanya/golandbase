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
	i              = 0
	limit          = 160
	// errType        string
	errExcessLimit = "limit error"
	errEOF         = errors.New("fuck EOF error")
	// errAnother     = errors.New("some another error")
)

func main() {
	err := readFile(pathin)
	if err != nil {
		if err == err.(*internal.ErrParser) {
			// errType = errEOF
			fmt.Println("string count exceed limit, please read another file =) err: ", err.Error())
		}
		if err == errEOF {
			// errType = errEOF
			fmt.Println(errEOF)
		} else {
			// errType = errAnother
			fmt.Println(err)
		}
	}
	// switch finalParse{
	// case finalParse.(*internal.ErrParser):
	// 	fmt.Println("string count exceed limit, please read another file =) err: ", finalParse.Error())
	// case errEOF:
	// 	fmt.Println(errEOF)
    // case nil:
	// 	fmt.Println("No error")
	// default:
	// 	fmt.Println(finalParse)
	// }
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
			// errType = errExcessLimit
			mylimit := internal.ErrParser{
				errExcessLimit,
				limit,
				value}
            // result = startParse(mylimit)
			// return errors.New(result)
			return &mylimit
		}
		if err != nil {
			if err == io.EOF {
				// errType = errEOF
				return errEOF
			} else {
				// errType = errAnother
				return err
			}
		}
	}
	return nil
}

// func startParse(mylimit internal.ErrParser) string {
// 	return mylimit.Error()
// }

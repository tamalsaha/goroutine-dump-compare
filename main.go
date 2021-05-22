package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	f1 = "/home/tamal/Downloads/goroutine1.txt"
	f2 = "/home/tamal/Downloads/goroutine2.txt"
)

type Diff struct {
	A string
	B string
}

func main() {
	l1, err := process(f1)
	if err != nil {
		log.Fatalln(err)
	}
	l2, err := process(f2)
	if err != nil {
		log.Fatalln(err)
	}

	cmp := map[int]*Diff{}
	for k, v := range l1 {
		d := cmp[k]
		if d == nil {
			d = new(Diff)
		}
		d.A = v
		cmp[k] = d
	}
	for k, v := range l2 {
		d := cmp[k]
		if d == nil {
			d = new(Diff)
		}
		d.B = v
		cmp[k] = d
	}

	for k, d := range cmp {
		if d.A == "" && d.B != "" {
			fmt.Println("Goroutine", k)
			fmt.Println(d.B)
			fmt.Println("--------------------------------------------------------")
		}
	}
}

// https://www.geeksforgeeks.org/how-to-read-a-file-line-by-line-to-string-in-golang/
func process(filename string) (map[int]string, error) {
	rts := map[int]string{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)

	rtNum := 0
	var st []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			rts[rtNum] = strings.Join(st, "\n")
			continue
		}

		if strings.HasPrefix(line, "goroutine ") {
			parts := strings.Split(line, " ")
			if len(parts) >= 2 {
				n, _ := strconv.Atoi(parts[1])
				rtNum = n
			}
			st = nil
		} else {
			st = append(st, line)
		}
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	return rts, nil
}

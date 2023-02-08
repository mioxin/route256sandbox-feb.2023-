package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := sum2(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}

}
func sum2(in io.Reader, out io.Writer) error {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var f int
	fmt.Fscan(r, &f)
	for i := 0; i < f; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		//fmt.Println(a, b)
		fmt.Fprintln(w, a+b)
	}
	return nil
}

func sum(in io.Reader, out io.Writer) error {
	r := bufio.NewScanner(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	if ok := r.Scan(); !ok {
		return fmt.Errorf("end of input")
	}
	f, err := strconv.Atoi(r.Text())
	if err != nil {
		return fmt.Errorf("first line is not int. %v\n", err)
	}
	for i := 0; i < f; i++ {
		if ok := r.Scan(); !ok {
			return fmt.Errorf("end of input")
		}
		af := strings.Split(r.Text(), " ")

		if len(af) < 2 {
			return fmt.Errorf("expectes 2 numbers")
		}

		f1, err := strconv.Atoi(af[0])
		f2, err1 := strconv.Atoi(af[1])
		if err != nil || err1 != nil {
			return fmt.Errorf("the line is not int. %v\n", err)
		}
		w.WriteString(fmt.Sprintf("%v\n", f1+f2))
	}
	return nil
}

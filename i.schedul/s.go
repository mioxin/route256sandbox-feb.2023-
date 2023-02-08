package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// fin, _ := os.Open("./tests/01")
	// err := run(fin, os.Stdout)
	err := run(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}


func get_times(r *bufio.Reader, t int) chan string {
	s_times := make(chan string)
	go func(ch chan string) {
		for i := 0; i < t; i++ {
			str_n, _ := r.ReadString('\n')
			ch <- str_n
		}
		close(ch)
	}(s_times)
	return s_times
}

func run(in io.Reader, out io.Writer) error {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	str, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	p, t, err := get_pair(str)
	if err != nil {
		return err
	}

	str, err = r.ReadString('\n')
	if err != nil {
		return err
	}
	procs := make([]int, p)
	for i, s := range strings.Split(strings.Trim(str, "\r\n"), " ") {
		procs[i], err = strconv.Atoi(s)
		if err != nil {
			return err
		}
	}
	sort.Ints(procs)

	sum := 0
	tproc := make(map[int]int, p)
	ar_times := make([]string, 0, 5000000)
	for i := 0; i < t; i++ {
		str_n, _ := r.ReadString('\n')
		ar_times = append(ar_times, str_n)
	}

	num_proc := 0
	for _, str_n := range ar_times {
		ts, tl, err := get_pair(str_n)
		if err != nil {
			return err
		}

		for i := num_proc; i < p; i++ {
			if tproc[i] < ts {
				tproc[i] = ts + tl - 1
				sum += procs[i] * tl
				break
			}
		}
		num_proc = 0
	}
	// for i, v := range tproc {
	// 	sum += v[1] * procs[i]
	// }
	//fmt.Println(j, ": ", outpu)
	w.WriteString(strconv.Itoa(sum))
	w.WriteString("\n")

	return nil
}

func get_pair(str string) (int, int, error) {
	a_str := strings.Split(strings.Trim(str, "\r\n"), " ")
	n, err := strconv.Atoi(a_str[0])
	if err != nil {
		return 0, 0, fmt.Errorf("first line is not int.\n %v", err)
	}
	m, err := strconv.Atoi(a_str[1])
	if err != nil {
		return 0, 0, fmt.Errorf("first line is not int.\n %v", err)
	}
	return n, m, nil
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type interval struct {
	t1, t2 time.Time
}

var arr_int []interval

func main() {
	// fin, _ := os.Open("./tests/01")
	// err := work(fin, os.Stdout)
	err := work(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}

func work(in io.Reader, out io.Writer) error {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	str_t, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	t, err := strconv.Atoi(strings.Trim(str_t, "\r\n"))
	if err != nil {
		return fmt.Errorf("first line is not int.\n %v", err)
	}

	for i := 0; i < t; i++ {
		okparse := true
		str_n, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		n, err := strconv.Atoi(strings.Trim(str_n, "\r\n"))
		//fmt.Print(">>> str_n: ", n, str_n)
		if err != nil {
			return fmt.Errorf("2nd line is not int.\n %v", err)
		}
		for i := 0; i < n; i++ {
			str, err := r.ReadString('\n')
			if err != nil {
				return err
			}
			arr := strings.Split(strings.Trim(str, "\r\n"), "-")
			s1, s2, err := get_time(&arr)
			if err != nil {
				okparse = false
			} else {
				arr_int = append(arr_int, interval{s1, s2})
			}
		}
		if !okparse {
			w.WriteString("NO\n")
		} else if n == 1 {
			w.WriteString("YES\n")
		} else if is_ok(&arr_int) {
			w.WriteString("YES\n")
		} else {
			w.WriteString("NO\n")
		}

		arr_int = nil
	}
	return nil
}

func get_time(astr *[]string) (time.Time, time.Time, error) {
	s1, err1 := time.Parse("15:04:05", (*astr)[0])
	s2, err2 := time.Parse("15:04:05", (*astr)[1])
	if err1 != nil {
		return s1, s2, fmt.Errorf("can't parse 1st time %v\n%v", (*astr)[0], err1)
	}
	if err2 != nil {
		return s1, s2, fmt.Errorf("can't parse 2nd time %v\n%v", (*astr)[0], err2)
	}
	if s1.After(s2) {
		return s1, s2, fmt.Errorf("%v must be befor %v", s1, s2)
	}
	return s1, s2, nil
}

func is_ok(arr *[]interval) bool {
	for i, it1 := range *arr {
		for _, it2 := range (*arr)[i+1:] {
			if (it1.t2.After(it2.t1) || it1.t2.Equal(it2.t1)) && (it2.t2.After(it1.t1) || it2.t2.Equal(it1.t1)) {
				return false
			}
		}
	}
	return true
}

func is_ok2(arr *[]interval) bool {
	sort.SliceStable(*arr, func(i, j int) bool {
		return (*arr)[i].t1.Before((*arr)[j].t1)
	})
	for idx, intrv := range *arr {
		if idx+1 >= len(*arr) {
			break
		}
		if intrv.t2.After((*arr)[idx+1].t1) || intrv.t2.Equal((*arr)[idx+1].t1) {
			return false
		}
	}
	return true
}

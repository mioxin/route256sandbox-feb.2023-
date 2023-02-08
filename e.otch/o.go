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
		str_n, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		n, err := strconv.Atoi(strings.Trim(str_n, "\r\n"))
		if err != nil {
			return fmt.Errorf("first line is not int.\n %v", err)
		}

		af, err := get_int_array(r, n)
		if err != nil {
			return err
		}

		if is_ok2(af) {
			w.WriteString("YES\n")
		} else {
			w.WriteString("NO\n")
		}
	}
	return nil
}

func contain[T int | string](a *[]T, memb T) bool {
	ok := false
	for _, m := range *a {
		if memb == m {
			ok = true
			break
		}
	}
	return ok
}

func is_ok2[T int | string](arr *[]T) bool {
	ok := true
	cache := make(map[T]int)
	last := (*arr)[0]
	cache[last]++
	length := len(*arr)
	for i := 1; i < length; i++ {
		if (*arr)[i] == last {
			continue
		}
		if _, is := cache[(*arr)[i]]; is {
			ok = false
			break
		}
		cache[(*arr)[i]] = 1
		last = (*arr)[i]
	}
	return ok
}

func is_ok[T int | string](arr *[]T) bool {
	ok := true
	cache := []T{}
	cache = append(cache, (*arr)[0])
	length := len(*arr)
	for i := 1; i < length; i++ {
		if (*arr)[i] == (*arr)[i-1] {
			continue
		}
		if contain(&cache, (*arr)[i]) {
			ok = false
			break
		}
		cache = append(cache, (*arr)[i])
	}
	return ok
}

func get_int_array(r *bufio.Reader, count int) (*[]int, error) {
	str, err := r.ReadString('\n')
	var arr_int []int
	if err != nil {
		return &arr_int, err
	}
	arr_str := strings.Split(strings.Trim(str, "\r\n"), " ")
	//fmt.Println(arr_str)
	if count > 0 && len(arr_str) != count {
		return &arr_int, fmt.Errorf("expectes %d numbers: %v", count, str)
	}
	for _, s := range arr_str {
		res, err := strconv.Atoi(s)
		if err != nil {
			return &arr_int, err
		}
		arr_int = append(arr_int, res)
	}
	return &arr_int, nil
}

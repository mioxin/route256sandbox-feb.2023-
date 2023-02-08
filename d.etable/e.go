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
		_, err := r.ReadString('\n')
		if err != nil {
			return err
		}

		af, err := get_int_array(r, 2)
		if err != nil {
			return err
		}

		n, m := (*af)[0], (*af)[1]
		etab := [][]int{}
		for i := 0; i < n; i++ {
			af, err = get_int_array(r, m)
			if err != nil {
				return err
			}
			etab = append(etab, *af)
		}

		af, err = get_int_array(r, 1)
		if err != nil {
			return err
		}

		clicks, err := get_int_array(r, (*af)[0])

		if err != nil {
			return err
		}
		for _, cl := range *clicks {
			sorting(&etab, cl-1)
		}
		for _, e := range etab {
			w.WriteString(strings.Trim(fmt.Sprintf("%v", e), "[]"))
			w.WriteString("\n")
		}
		w.WriteString("\n")
	}
	return nil
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

func sorting(a *[][]int, field int) {
	sort.SliceStable(*a, func(i, j int) bool {

		return (*a)[i][field] < (*a)[j][field]
	})

}

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

var users map[int]map[int]int

func main() {
	// fin, _ := os.Open("./tests/12")
	// err := run(fin, os.Stdout)
	err := run(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}

func run(in io.Reader, out io.Writer) error {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	str, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	n, m, err := get_pair(str)
	if err != nil {
		return err
	}

	users = make(map[int]map[int]int, n)
	for i := 0; i < m; i++ {
		str_n, _ := r.ReadString('\n')
		x, y, err := get_pair(str_n)
		if err != nil {
			return err
		}

		friends, inMap := users[x]
		if inMap {
			friends[y] = 0
		} else {
			friends := make(map[int]int, 5)
			friends[y] = 0
			users[x] = friends
		}

		friends1, inMap := users[y]
		if inMap {
			friends1[x] = 0
		} else {
			friends1 := make(map[int]int, 5)
			friends1[x] = 0
			users[y] = friends1
		}
	}
	// for k := 1; k <= n; k++ {
	// 	fmt.Println(k, ":", users[k])
	// }
	for j := 1; j <= n; j++ {
		outpu := get_friend(j)
		if len(outpu) == 0 {
			//w.WriteString(strconv.Itoa(j))
			w.WriteString("0\n")
		} else {
			//fmt.Println(j, ": ", outpu)
			w.WriteString(get_max_value(outpu))
			w.WriteString("\n")

		}
	}

	return nil
}

func get_friend(u int) map[int]int {
	var r map[int]int = nil
	friends, inMap := users[u]
	if !inMap {
		return r
	}
	r = make(map[int]int)
	for f, _ := range friends {
		for f2, _ := range users[f] {
			_, inMap := users[u][f2]
			if f2 != u && !inMap {
				r[f2]++
			}
		}
	}
	return r
}

func get_max_value(f map[int]int) string {
	maxv := 0
	i_f := make([]int, 0, 5)
	for f, v := range f {
		if v > maxv {
			maxv = v
			i_f = nil
			i_f = append(i_f, f)
		} else if v == maxv {
			i_f = append(i_f, f)
		}
	}
	sort.Ints(i_f)

	//return strings.Trim(fmt.Sprint(i_f), "[]")
	return int_slice_to_string(i_f)
}

func int_slice_to_string(is []int) string {
	r := ""
	if len(is) == 1 {
		return strconv.Itoa(is[0])
	}
	for _, i := range is {
		r = r + " " + strconv.Itoa(i)
	}
	r = strings.Trim(r, " ")
	return r
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

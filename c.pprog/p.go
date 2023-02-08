package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fin, _ := os.Open("./tests/06")
	// err := pairs(fin, os.Stdout)
	err := pairs(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
	//	init_cache(5)
}

func init_cache(n int) *[][]int {
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
		for j := 0; j < n; j++ {
			r[i][j] = 101
		}
	}
	return &r
}

func pairs(in io.Reader, out io.Writer) error {
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
		str_f, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		f, _ := strconv.Atoi(strings.Trim(str_f, "\r\n"))
		cache := init_cache(f)

		str_af, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		af := strings.Split(strings.Trim(str_af, "\r\n"), " ")

		if len(af) != f {
			return fmt.Errorf("expectes %d numbers: %v", f, af)
		}
		devs := make([]int, f)
		for ind, s := range af {
			ns, _ := strconv.Atoi(s)
			devs[ind] = ns
		}
		for i := 0; i < f/2-1; i++ {
			dev1, dev2 := get_pair(&devs, cache)
			w.WriteString(fmt.Sprintf("%d %d\n", dev1, dev2))
		}
		//осталась единственная пара
		p := [2]int{}
		i := 0
		for ind, d := range devs {
			if d != 0 {
				p[i] = ind + 1
				i++
			}
			if i > 1 {
				break
			}
		}
		w.WriteString(fmt.Sprintf("%d %d\n\n", p[0], p[1]))
	}
	return nil
}

func get_pair(d *[]int, c *[][]int) (int, int) {
	min := 100
	m := 0
	m_d1 := 0
	d1, d2 := 0, 0
	for i, dev := range *d {
		if dev == 0 {
			continue
		}
		d1 = i
		m_d1 = dev
		(*d)[i] = 0
		for j, dev := range *d {
			if dev == 0 {
				continue
			}
			if (*c)[d1][j] < 101 {
				m = (*c)[d1][j]
			} else {
				m = int(math.Abs(float64(m_d1) - float64(dev)))
				(*c)[d1][j] = m
			}
			if m == 0 {
				d2 = j
				(*d)[j] = 0

				break
			}
			if m < min {
				d2 = j
				min = m
			}
		}
		(*d)[d2] = 0
		break

	}
	return d1 + 1, d2 + 1
}

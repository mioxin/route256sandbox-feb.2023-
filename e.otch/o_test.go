package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

type test_cont struct {
	ar   []int
	ch   int
	want bool
}

var t1 = []test_cont{
	{[]int{1, 2, 3, 4}, 2, true},
	{[]int{1, 2, 3, 4}, 10, false},
}

type test_isok struct {
	ar   []int
	want bool
}

var t2 = []test_isok{
	{[]int{1, 2, 2, 2, 5, 7}, true},
	{[]int{1}, true},
	{[]int{2, 2, 2, 2, 2, 2}, true},
	{[]int{1, 2, 2, 2, 5, 2}, false},
	{[]int{2, 7, 2, 2, 2, 2}, false},
}

type test struct {
	inf, resf string
}

var tests []test

func TestMain(m *testing.M) {
	for i := 1; i < 10; i++ {
		tests = append(tests, test{"./tests/0" + strconv.Itoa(i), "./tests/0" + strconv.Itoa(i) + ".a"})
	}
	for i := 10; i < 41; i++ {
		tests = append(tests, test{"./tests/" + strconv.Itoa(i), "./tests/" + strconv.Itoa(i) + ".a"})
	}
	m.Run()
}

func Test_contain(t *testing.T) {
	for _, te := range t1 {
		t.Run(fmt.Sprintf("%v", te.ar), func(t *testing.T) {
			in := te.ar
			if contain(&in, te.ch) != te.want {
				t.Errorf("shold be %v", te.want)
			}
		})
	}
}

func Test_isok(t *testing.T) {
	for _, te := range t2 {
		t.Run(fmt.Sprintf("%v", te.ar), func(t *testing.T) {
			in := te.ar
			if is_ok(&in) != te.want {
				t.Errorf("shold be %v but %v", te.want, is_ok(&in))
			}
		})
	}
}

func Test_sort(t *testing.T) {
	for _, ts := range tests {
		fmt.Println("test:", ts.inf, " ", ts.resf)
		fin, err := os.Open(ts.inf)
		if err != nil {
			fmt.Println(err)
		}
		defer fin.Close()
		os.Remove("./out")
		fout, err := os.OpenFile("./out", os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			t.Error(err)
		}
		defer fout.Close()

		fr, err := os.Open(ts.resf)
		if err != nil {
			fmt.Println(err)
		}
		defer fr.Close()

		if err := work(fin, fout); err != nil {
			fmt.Println(err)
		}

		fout.Close()
		fout, err = os.Open("./out")
		if err != nil {
			t.Error(err)
		}
		fr.Close()
		fr, _ = os.Open(ts.resf)

		in1 := bufio.NewReader(fout)
		in2 := bufio.NewReader(fr)
		for {
			s1, err := in1.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Errorf("error scan in1\n%v", err)
				break
			}
			s2, err := in2.ReadString('\n')
			if err != nil {
				t.Errorf("error scan in2\n%v", err)
				break
			}
			//fmt.Println(">>>", s1, " ", s2)
			if strings.Trim(s1, "\r\n") == strings.Trim(s2, "\r\n") {
				//fmt.Printf("ok>> %v//%v :%v\n", s1, s2, s1 == s2)
			} else {
				t.Errorf("%v//%v :%v", s1, s2, s1 == s2)
			}
		}
		fin.Close()
		fout.Close()
		fr.Close()
	}

}

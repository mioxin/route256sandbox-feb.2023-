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

var get = [][]int{
	{3, 4, 1},
	{2, 2, 5},
	{2, 4, 2},
	{2, 2, 1},
}

var want = [][]int{
	{2, 2, 1},
	{3, 4, 1},
	{2, 4, 2},
	{2, 2, 5},
}

var click = []int{2, 1, 3}

type test struct {
	inf, resf string
}

var tests []test

func TestMain(m *testing.M) {
	for i := 1; i < 10; i++ {
		tests = append(tests, test{"./tests/0" + strconv.Itoa(i), "./tests/0" + strconv.Itoa(i) + ".a"})
	}
	for i := 10; i < 11; i++ {
		tests = append(tests, test{"./tests/" + strconv.Itoa(i), "./tests/" + strconv.Itoa(i) + ".a"})
	}
	m.Run()
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
				t.Errorf("\n%v\n%v :%v\n", s1, s2, s1 == s2)
			}
		}
		fin.Close()
		fout.Close()
		fr.Close()
	}

}
func Test_sorting(t *testing.T) {
	for _, c := range click {
		sorting(&get, c-1)
	}
	for i, r := range get {
		//fmt.Println(r, " // ", want[i])
		for j, c := range r {
			if c != want[i][j] {
				t.Errorf("i:%v j:%v get/want: %v/%v\n", i, j, c, want[i][j])
			}
		}
	}
}

func Test_get_array(t *testing.T) {
	str_get := "1 2 3 4\n"
	want := []int{1, 2, 3, 4}
	s := bufio.NewReader(strings.NewReader(str_get))
	arr, err := get_int_array(s, 4)
	if err != nil {
		t.Error(*arr, err)
	}
	for idx, a := range *arr {
		if want[idx] != a {
			t.Error(*arr, " != ", want)
		}
	}
	fmt.Printf("%v", *arr)
}

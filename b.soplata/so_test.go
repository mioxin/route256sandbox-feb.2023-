package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

type test struct {
	inf, resf string
}

var tests []test

func TestMain(m *testing.M) {
	for i := 1; i < 10; i++ {
		tests = append(tests, test{"./tests/0" + strconv.Itoa(i), "./tests/0" + strconv.Itoa(i) + ".a"})
	}
	tests = append(tests, test{"./tests/10", "./tests/10.a"})
	m.Run()
}

func Test_so(t *testing.T) {
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

		if err := so(fin, fout); err != nil {
			fmt.Println(err)
		}
		fout, err = os.Open("./out")
		if err != nil {
			t.Error(err)
		}

		in1 := bufio.NewScanner(fout)
		in2 := bufio.NewScanner(fr)
		for {
			ok1 := in1.Scan()
			if !ok1 {
				fmt.Println(in1.Err(), in2.Err())
				t.Errorf("%v//%v :error scan in1\n", in1.Err(), in2.Err())
				break
			}
			if !in2.Scan() {
				fmt.Println(in1.Err(), in2.Err())
				t.Errorf("%v//%v :error scan in2\n", in1.Err(), in2.Err())
				break
			}
			s1 := in1.Text()
			s2 := in2.Text()
			//fmt.Println(">>>", s1, " ", s2)
			if s1 == s2 {
				//fmt.Printf("ok>> %v//%v :%v\n", s1, s2, s1 == s2)
			} else {
				t.Errorf("%v//%v :%v\n", s1, s2, s1 == s2)
			}
		}
		fin.Close()
		fout.Close()
		fr.Close()
	}

}

func Benchmark_so(b *testing.B) {

}

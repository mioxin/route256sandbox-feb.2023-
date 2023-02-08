package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var tests = []struct {
	inf, resf string
}{
	{"./01", "./01.a"},
	{"./02", "./02.a"},
}

// func MainTest(m *testing.M){

// }
func Test_sum(t *testing.T) {
	for _, test := range tests {
		fmt.Println("test:", test.inf, " ", test.resf)
		fin, err := os.Open(test.inf)
		if err != nil {
			fmt.Println(err)
		}
		defer fin.Close()

		fout, err := os.OpenFile("./out", os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			t.Error(err)
		}
		defer fout.Close()

		fr, err := os.Open(test.resf)
		if err != nil {
			fmt.Println(err)
		}
		defer fr.Close()

		if err := sum2(fin, fout); err != nil {
			fmt.Println(err)
		}
		fout, err = os.OpenFile("./out", os.O_RDONLY, 0777)
		if err != nil {
			t.Error(err)
		}

		in1 := bufio.NewScanner(fout)
		in2 := bufio.NewScanner(fr)
		for in1.Scan() {
			if !in2.Scan() {
				fmt.Println("Can't read file rez")
				break
			}
			s1 := in1.Text()
			s2 := in2.Text()
			fmt.Println(">>>", s1, " ", s2)
			if s1 == s2 {
				fmt.Printf("or>> %v//%v :%v\n", s1, s2, s1 == s2)
			} else {
				t.Errorf("%v//%v :%v\n", s1, s2, s1 == s2)
			}
		}
		fin.Close()
		fout.Close()
		fr.Close()
	}
}

func Benchmark_sum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//fmt.Println("test:", test.inf[0], " ", test.resf[0])
		fin, err := os.Open("./02")
		if err != nil {
			fmt.Println(err)
		}
		defer fin.Close()

		fout, err := os.OpenFile("./out", os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			b.Error(err)
		}
		defer fout.Close()

		fr, err := os.Open("./02.a")
		if err != nil {
			fmt.Println(err)
		}
		defer fr.Close()

		sum(fin, fout)
	}

}
func Benchmark_sum2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//fmt.Println("test:", test.inf[0], " ", test.resf[0])
		fin, err := os.Open("./02")
		if err != nil {
			fmt.Println(err)
		}
		defer fin.Close()

		fout, err := os.OpenFile("./out", os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			b.Error(err)
		}
		defer fout.Close()

		fr, err := os.Open("./02.a")
		if err != nil {
			fmt.Println(err)
		}
		defer fr.Close()

		sum2(fin, fout)
	}

}

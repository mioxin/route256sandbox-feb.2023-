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

type test struct {
	inf, resf string
}

var tests []test

type develop struct {
	f      int
	dev    string
	result []string
}

var dev = develop{
	6,
	"2 3 2 4 1 3",
	[]string{"1 3", "2 6", "4 5"},
}

func TestMain(m *testing.M) {
	for i := 1; i < 10; i++ {
		tests = append(tests, test{"./tests/0" + strconv.Itoa(i), "./tests/0" + strconv.Itoa(i) + ".a"})
	}
	for i := 10; i < 21; i++ {
		tests = append(tests, test{"./tests/" + strconv.Itoa(i), "./tests/" + strconv.Itoa(i) + ".a"})
	}
	m.Run()
}

func Test_getpair(t *testing.T) {
	cache := init_cache(dev.f)
	d := strings.Split(dev.dev, " ")
	devs := make([]int, dev.f)
	for ind, s := range d {
		ns, _ := strconv.Atoi(s)
		devs[ind] = ns
	}

	for i := 0; i < dev.f/2; i++ {
		d1, d2 := get_pair(&devs, cache)
		fmt.Printf("%v %v -- %v\n", d1, d2, dev.result[i])
		if fmt.Sprintf("%v %v", d1, d2) != dev.result[i] {
			t.Errorf("%v %v != %v", d1, d2, dev.result[i])
		}
	}
}

func Test_pair(t *testing.T) {
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

		if err := pairs(fin, fout); err != nil {
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
			if err == io.EOF {
				break
			}
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

func Benchmark_so(b *testing.B) {

}

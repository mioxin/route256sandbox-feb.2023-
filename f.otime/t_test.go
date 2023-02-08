package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type test_t struct {
	ar   []string
	want bool
}

var t1 = []test_t{
	{[]string{"02:46:00", "03:14:59"}, true},
	{[]string{"04:46:00", "03:14:59"}, false},
	{[]string{"02:66:00", "03:14:59"}, false},
	{[]string{"02:26:00", "24:14:59"}, false},
	{[]string{"02:46:00", "02:46:00"}, true},
}
var t_ok = []test_t{
	{[]string{"02:46:00", "03:14:59"}, true},
	{[]string{"03:46:00", "04:46:00"}, true},
}
var t_nok = []test_t{
	{[]string{"02:46:00", "03:14:59"}, true},
	{[]string{"03:14:59", "04:46:00"}, true},
}
var t_nok2 = []test_t{
	{[]string{"02:46:00", "03:14:59"}, true},
	{[]string{"03:14:59", "04:46:00"}, true},
}

type test struct {
	inf, resf string
}

var tests []test

func TestMain(m *testing.M) {
	for i := 1; i < 10; i++ {
		tests = append(tests, test{"./tests/0" + strconv.Itoa(i), "./tests/0" + strconv.Itoa(i) + ".a"})
	}
	for i := 10; i < 36; i++ {
		tests = append(tests, test{"./tests/" + strconv.Itoa(i), "./tests/" + strconv.Itoa(i) + ".a"})
	}
	m.Run()
}

func Test_get_t(t *testing.T) {
	for _, te := range t1 {
		t.Run(fmt.Sprintf("%v", te.ar), func(t *testing.T) {
			in := te.ar
			if _, _, err := get_time(&in); err != nil {
				if te.want {
					t.Errorf("shold be %v, but %v", te.want, err)
				}
			} else if !te.want {
				t.Errorf("shold be %v, but %v", te.want, err)
			}
		})
	}
}

func Test_isok(t *testing.T) {
	var arr_int []interval
	for _, sct := range t_nok2 {
		s1, _ := time.Parse("15:04:05", sct.ar[0])
		s2, _ := time.Parse("15:04:05", sct.ar[1])
		arr_int = append(arr_int, interval{s1, s2})
	}
	if is_ok(&arr_int) {
		t.Error("err")
	}
}
func Test_work(t *testing.T) {
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

func Benchmark_work(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//fmt.Println("test:", test.inf[0], " ", test.resf[0])
		fin, err := os.Open("./tests/34")
		if err != nil {
			fmt.Println(err)
		}
		defer fin.Close()

		fout, err := os.OpenFile("./out", os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			b.Error(err)
		}
		defer fout.Close()

		work(fin, fout)
	}

}

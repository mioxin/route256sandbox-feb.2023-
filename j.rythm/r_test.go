package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

type test_t struct {
	ar   []string
	want bool
}

type test struct {
	inf, resf string
}

var tests []test

func TestMain(m *testing.M) {
	for i := 1; i < 10; i++ {
		tests = append(tests, test{"./tests/0" + strconv.Itoa(i), "./tests/0" + strconv.Itoa(i) + ".a"})
	}
	for i := 10; i < 26; i++ {
		tests = append(tests, test{"./tests/" + strconv.Itoa(i), "./tests/" + strconv.Itoa(i) + ".a"})
	}
	m.Run()
}

func Test_run(t *testing.T) {
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

		if err := run(fin, fout); err != nil {
			fmt.Println(err)
		}

		fin.Close()
		fin, err = os.Open(ts.inf)
		if err != nil {
			t.Error(err)
		}

		in := bufio.NewReader(fin)
		var words *[]string // массив слов для поиска
		_, err = get_arr(in)
		if err != nil {
			t.Errorf("error scan ts.inf\n%v", err)
		}

		words, err = get_arr(in)
		if err != nil {
			t.Errorf("error scan ts.inf\n%v", err)
		}

		fout.Close()
		fout, err = os.Open("./out")
		if err != nil {
			t.Error(err)
		}
		fr.Close()
		fr, err = os.Open(ts.resf)
		if err != nil {
			t.Error(err)
		}

		stream_out := bufio.NewReader(fout)
		stream_res := bufio.NewReader(fr)
		i := 0
		for {
			s_out, err := stream_out.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Errorf("error scan stream_out\n%v", err)
				break
			}
			s_res, err := stream_res.ReadString('\n')
			if err != nil {
				t.Errorf("error scan stream_res\n%v", err)
				break
			}
			//fmt.Println(">>>", s_out, " ", s_res)
			s_out = strings.Trim(s_out, "\r\n")
			s_res = strings.Trim(s_res, "\r\n")
			l_res := level_r((*words)[i], s_res)
			l_out := level_r((*words)[i], s_out)
			if (s_out != s_res && l_out < l_res) || s_out == (*words)[i] {
				t.Errorf("%v: \t%v//%v :%v lev:%v/%v", (*words)[i], s_out, s_res, s_out == s_res, l_out, l_res)
			}
			//fmt.Printf("ok>> запрос %v: %v//%v :%v\n", (*words)[i], s_out, s_res, s_out == s_res)
			i++
		}
		fin.Close()
		fout.Close()
		fr.Close()
	}

}
func Test_run2(t *testing.T) {
	inf := "./tests/01"
	resf := "./tests/01.a"
	fmt.Println("test:", inf, " ", resf)
	fin, err := os.Open(inf)
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

	fr, err := os.Open(resf)
	if err != nil {
		fmt.Println(err)
	}
	defer fr.Close()

	if err := run(fin, fout); err != nil {
		fmt.Println(err)
	}

	fin.Close()
	fin, err = os.Open(inf)
	if err != nil {
		t.Error(err)
	}

	in := bufio.NewReader(fin)
	var words *[]string // массив слов для поиска
	_, err = get_arr(in)
	if err != nil {
		t.Errorf("error scan ts.inf\n%v", err)
	}

	words, err = get_arr(in)
	if err != nil {
		t.Errorf("error scan ts.inf\n%v", err)
	}

	fout.Close()
	fout, err = os.Open("./out")
	if err != nil {
		t.Error(err)
	}
	fr.Close()
	fr, err = os.Open(resf)
	if err != nil {
		t.Error(err)
	}

	stream_out := bufio.NewReader(fout)
	stream_res := bufio.NewReader(fr)
	i := 0
	for {
		s_out, err := stream_out.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Errorf("error scan stream_out\n%v", err)
			break
		}
		s_res, err := stream_res.ReadString('\n')
		if err != nil {
			t.Errorf("error scan stream_res\n%v", err)
			break
		}
		//fmt.Println(">>>", s_out, " ", s_res)
		s_out = strings.Trim(s_out, "\r\n")
		s_res = strings.Trim(s_res, "\r\n")
		l_res := level_r((*words)[i], s_res)
		l_out := level_r((*words)[i], s_out)
		if (s_out != s_res && l_out < l_res) || s_out == (*words)[i] {
			t.Errorf("%v: \t%v//%v :%v lev:%v/%v", (*words)[i], s_out, s_res, s_out == s_res, l_out, l_res)
		}
		// if s_out != s_res {
		// 	fmt.Printf("ok>> запрос %v: %v//%v :%v lev:%v/%v\n", (*words)[i], s_out, s_res, s_out == s_res, l_out, l_res)
		// }

		i++
	}
	fin.Close()
	fout.Close()
	fr.Close()

}

func level_r(w, r string) int {
	i := 0
	l := int(math.Min(float64(len(w)), float64(len(r))))
	w = reverse(w)
	r = reverse(r)
	for i = 0; i < l; i++ {
		if w[i] != r[i] {
			break
		}
	}
	return i
}
func Test_srt(t *testing.T) {
	s := []string{
		"pppjs",
		"jpjpsp",
		"s",
		"ppsssjjj",
		"sppjjjps",
		"jjsjsppps",
		"jpjjpss",
		"jppjspspjj",
		"sssjjjp",
		"p",
		"spppj",
		"sjjjjpjjss",
		"jsjjjsss",
		"sjss",
		"jpjj",
		"ppj",
		"spss",
		"sjp",
		"jjjpppj",
		"sjppjssj",
	}
	srt(s)
}
func Benchmark_work(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//fmt.Println("test:", test.ts.inf[0], " ", test.resf[0])
		fin, err := os.Open("./tests/25")
		if err != nil {
			fmt.Println(err)
		}
		defer fin.Close()

		fout, err := os.OpenFile("./out", os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			b.Error(err)
		}
		defer fout.Close()

		run(fin, fout)
	}

}

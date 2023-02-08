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

type pair_int struct {
	i, j int
}

func main() {
	fin, _ := os.Open("./tests/00")
	err := run(fin, os.Stdout)
	//	err := run(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}

func run(in io.Reader, out io.Writer) error {
	n_pool := 13
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	dict := NewDict(r)
	//full_index := dict.index
	words, err := get_arr(r)
	if err != nil {
		return err
	}

	quit := make(chan struct{})
	output := make(chan string)

	go func(quit chan struct{}, o chan string) {
		pool := make(chan int, n_pool)
		for i := 0; i < n_pool; i++ {
			pool <- i + 1
		}

		start := make(chan struct{})
		//done := make(chan struct{})
		go func(s chan struct{}) {
			s <- struct{}{}
		}(start)

		for _, word := range *words {
			id := <-pool
			//fmt.Println("GGGGGGGGGGGGGGGGGo word:", word)

			start = dict.go_search(pool, id, word, start, o)
		}

		for i := 0; i < n_pool; i++ {
			<-pool
		}

		quit <- struct{}{}
	}(quit, output)

	go func(o chan string) {
		for s := range o {
			w.WriteString(s)
		}
	}(output)

	<-quit
	return nil
}

func (d *Dictionary) go_search(pool chan int, id int, w string, st chan struct{}, out1 chan string) chan struct{} {
	end_search := make(chan struct{})

	go func(start chan struct{}, end chan struct{}, w string) {
		//fmt.Println("GGGGGGGGGGGGGGGGGo_search id:", id)
		r := d.Get_Ryth4(w, pair_int{0, len(d.dict)})
		<-start
		out1 <- strconv.Itoa(id)
		out1 <- ": "
		out1 <- w
		out1 <- ":\t\t"
		out1 <- r
		out1 <- "\n"

		pool <- id
		end <- struct{}{}
		//fmt.Println("GGGGGGGGGGGGGGGGGo_END_search id:", id, r)

	}(st, end_search, w)

	return end_search
}

type Dictionary struct {
	dict    []string
	dict_sl [10][]byte
	index   map[byte]int
	//idx_pair pair_int
}

func NewDict(r *bufio.Reader) *Dictionary {
	arr, err := get_arr(r)
	if err != nil {
		panic(err)
	}
	d := new(Dictionary)
	d.dict = *arr
	for i, w := range d.dict {
		d.dict[i] = reverse(w)
	}
	sort.Strings(d.dict)

	sl := build_dict_slice(&d.dict)
	d.dict_sl = *sl

	d.index = make(map[byte]int, len(d.dict))
	for idx, ch := range d.dict {
		if _, inMap := d.index[ch[0]]; inMap {
			continue
		}
		d.index[ch[0]] = idx
	}
	//d_idx_pair = pair_int{0, len(d.dict)}
	return d
}

func build_dict_slice(dict *[]string) *[10][]byte {
	sl := [10][]byte{}
	for i := 0; i < 10; i++ {
		sl[i] = make([]byte, len(*dict))
	}
	for j, s := range *dict {
		for i := 0; i < 10; i++ {
			if len(s) < i+1 {
				sl[i][j] = 0
			} else {
				sl[i][j] = s[i]
			}
		}
	}
	return &sl
}

func reverse(s string) string {
	sb := []byte(s)
	for i, j := 0, len(sb)-1; i < j; i, j = i+1, j-1 {
		sb[i], sb[j] = sb[j], sb[i]
	}
	return string(sb)
}

func get_arr(r *bufio.Reader) (*[]string, error) {
	arr := make([]string, 0, 50000)
	str_t, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}
	d, err := strconv.Atoi(strings.Trim(str_t, "\r\n"))
	if err != nil {
		return nil, fmt.Errorf("first line %s is not int.\n %v", str_t, err)
	}
	//get dict
	for i := 0; i < d; i++ {
		str, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		arr = append(arr, (strings.Trim(str, "\r\n")))
		//////////fmt.Print(">>> str_n: ", n, str_n)
	}
	return &arr, nil
}

func (d *Dictionary) speed_search(sfull_revers, str string, isFirst bool, d_idx_pair pair_int) string {
	s := reverse(str)
	length := len(s)
	////fmt.Println(">>>", sfull_revers, s, isFirst)
	new_dict := d.dict[d.index[s[0]]:d_idx_pair.j]
	for _, sd := range new_dict {
		if isFirst && length < len(sd) && s == sd[:length] {
			////fmt.Println(">>>>>>1", sfull_revers, s, isFirst)

			return reverse(sd)
		} else if !isFirst && length <= len(sd) && s == sd[:length] && sfull_revers != sd {
			////fmt.Println(">>>>>>2", sfull_revers, s, isFirst)

			return reverse(sd)
		}
	}

	return ""
}

func srt(s []string) {
	sort.Strings(s)
	//sort.Search()
	// for _, str := range s {
	////fmt.Println(str)
	// }
}

type Dictionary2 struct {
	dict     []string
	index    map[byte]int
	idx_pair pair_int
}

func NewDict2(r *bufio.Reader) *Dictionary2 {
	arr, err := get_arr(r)
	if err != nil {
		panic(err)
	}
	d := new(Dictionary2)
	d.dict = *arr
	for i, w := range d.dict {
		d.dict[i] = reverse(w)
	}
	sort.Strings(d.dict)

	d.index = make(map[byte]int, len(d.dict))
	for idx, ch := range d.dict {
		if _, inMap := d.index[ch[0]]; inMap {
			continue
		}
		d.index[ch[0]] = idx
	}
	d.idx_pair = pair_int{0, len(d.dict)}
	return d
}

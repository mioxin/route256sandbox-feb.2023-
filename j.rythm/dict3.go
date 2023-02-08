package main

import (
	"bufio"
	"io"
)

func run3(in io.Reader, out io.Writer) error {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	dict := NewDict2(r)
	//full_index := dict.index
	words, err := get_arr(r)
	if err != nil {
		return err
	}

	for _, word := range *words {

		r := dict.Get_Ryth3(word)
		// w.WriteString(strconv.Itoa(id))
		// w.WriteString(": ")
		// w.WriteString(word)
		// w.WriteString(":\t\t")
		w.WriteString(r)
		w.WriteString("\n")
		//dict.idx_pair = pair_int{0, len(dict.dict)}
	}
	return nil
}

func (d *Dictionary2) Get_Ryth3(s string) string {
	s_rev := reverse(s)
	//sd := ""
	// //поиск с максимальной длинны
	// for i, _ := range s {
	// 	if i == 0 {
	// 		sd = d.speed_search(s_rev, s[i:], true)
	// 	} else {
	// 		sd = d.speed_search(s_rev, s[i:], false)
	// 	}
	// 	if sd != "" {
	//// 		////fmt.Println("speed search", sd)
	// 		return sd
	// 	}
	// }

	//поиск с одного символа

	//fmt.Printf("=====================================================\nNEW SEARCH")
	for i, ch := range reverse(s) {
		var inMap bool
		if i == 0 { //перый символ
			d.idx_pair.i, inMap = d.index[byte(ch)]
			if !inMap {
				break
			}
			if d.idx_pair.j-d.idx_pair.i == 1 {
				if d.dict[d.idx_pair.i][0] == byte(ch) && s_rev != d.dict[d.idx_pair.i] {
					return reverse(d.dict[d.idx_pair.i])
				} else {
					break
				}
			}
		}

		t_i, t_j := d.speed_search2(s_rev, s_rev[:i+1])
		//fmt.Printf("RYTH: t_i=%v, t_j=%v\n", t_i, t_j)
		if t_j-t_i == 1 { //осталась одна строка
			if s_rev != d.dict[t_i] {
				return reverse(d.dict[t_i])
			} else { // запрос равен единственной строке, пропускаем ее
				t_i++
			}
		}

		if t_i == t_j { //выбора больше нет, выводим из прошлого диапазона любую
			for _, ryth := range d.dict[d.idx_pair.i:d.idx_pair.j] {
				if ryth != s_rev {
					return reverse(ryth)
				}
			}
		}
		d.idx_pair.i, d.idx_pair.j = t_i, t_j
	}

	//рифм нет выводим любую
	for _, ds := range d.dict { //
		if ds[0] != s[len(s)-1] {
			return reverse(ds)
		}
	}

	return "!!!!!!!!!!!!!!"
}

func (d *Dictionary2) speed_search2(sfull_revers, s string) (int, int) {
	length := len(s)
	isFull := sfull_revers == s
	new_dict := d.dict[d.idx_pair.i:d.idx_pair.j]
	//fmt.Printf(">>>INP sfull_revers=%v, s=%v, new_dict:%v : %v (%v)\n", sfull_revers, s, d.idx_pair.i, d.idx_pair.j, length)
	var new_i, new_j int
	isFind := false
	sd := ""
	for new_i, sd = range new_dict {
		if length <= len(sd) {
			if isFull {
				if s == sd[:length] && length < len(sd) {
					//fmt.Printf(">>>> full find: sfull_revers=%v, s=%v > %v\n", sfull_revers, s, sd)
					return d.idx_pair.i + new_i, d.idx_pair.i + new_i + 1
				} else {
					continue
				}
			} else if s == sd[:length] {
				//fmt.Printf(">>>> find: sfull_revers=%v, s=%v > %v\n", sfull_revers, s, sd)
				isFind = true
				break
			}
		}
	}

	if !isFind { //весь массив пройден, рифма не найдена
		new_i++
		new_j = new_i
		//fmt.Printf(">>>> not find new_i=%v, sfull_revers=%v, s=%v\n", new_i, sfull_revers, s)
	} else { //ищем конец массива
		for new_j = new_i + 1; new_j < len(new_dict); new_j++ {
			if s != new_dict[new_j][:length] {
				break
			}
		}
		//fmt.Printf(">>>> END new_i=%v, new_j=%v, sfull_revers=%v, s=%v > %v\n", new_i, new_j, sfull_revers, s, new_dict[new_j-1:new_j])

	}
	return d.idx_pair.i + new_i, d.idx_pair.i + new_j
}

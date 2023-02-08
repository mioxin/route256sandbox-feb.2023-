package main

import (
	"bufio"
	"io"
)

func run2(in io.Reader, out io.Writer) error {
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

		r := dict.Get_Ryth2(word)
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

func (d *Dictionary2) NewPair(ch byte, count int, len_word int) *pair_int {
	// if len(d.dict) > 20 {
	////fmt.Println("===NEW_DICT_BY_CHAR===\n get from dict: ", ch, "; by char: ", string(ch), "; count: ", count)
	// } else {
	//fmt.Println("===NEW_DICT_BY_CHAR===\n get from dict: ", d, "; by char: ", string(ch), "; count: ", count)

	// }
	var inMap bool
	isLast := len_word == count+1
	var i_dict, j_dict int
	new_dict := d.dict[d.idx_pair.i:d.idx_pair.j]
	i_dict = -1    //индекс искомого символа
	if count > 0 { //не первоначальный массив (уже обработанный раньше)
		//строим индексы по заданному символу
		for idx, c := range new_dict {
			//if _, inMap := d.index[c[count]]; inMap {
			if c[count] == ch {
				i_dict = idx
				break
			}
			//fmt.Println("find index, ch/idx: ", c, c[count], idx)
			//d.index[c[count]] = idx
		}
	} else {
		i_dict, inMap = d.index[ch]
		if !inMap {
			//fmt.Println("not in map", reverse(string(new_dict[i_dict])), string(ch), ch, count, i_dict)
			////fmt.Println("not in map", d, reverse(string(new_dict[i_dict])), string(ch), ch, count, i_dict)
			return nil
		}
	}
	if i_dict < 0 {
		return nil
	}
	//d.idx_pair.i = i_dict
	for j_dict = i_dict; j_dict < len(new_dict); j_dict++ {
		//fmt.Println("idict", j_dict, ":", reverse(string(new_dict[j_dict])))
		//fmt.Println("есть рифма ? ch == d.dict[j][count]", ch, new_dict[j_dict][count])
		if ch == new_dict[j_dict][count] {
			//continue
		} else { //конец интервала в отсортированном массиве с нужным символом
			break
		}
	}
	//d.idx_pair.j = j_dict
	for i := i_dict; i < j_dict; i++ {
		//fmt.Printf("совпадение рифмы и запроса? %v. слово в словаре: %v (%v)>> %v\n", (isLast && (len_word == len(new_dict[i]))), reverse(string(new_dict[i])), len(new_dict[i]), d.dict[i_dict+d.idx_pair.i:j_dict+d.idx_pair.i])
		if !(isLast && (len_word == len(new_dict[i]))) { //последний символ + длинны запроса и слова совпадают - рифму не учитываем (по условию)
			//fmt.Printf("совпадение рифмы и запроса. i=%v, i_dict=%v, j_dict=%v\n", i, i_dict, j_dict)
			//continue
			//} else { //конец интервала в отсортированном массиве с нужным символом
			i_dict = i
			break
		}

	}
	if i_dict+1 == j_dict && isLast && (len_word == len(new_dict[i_dict])) { // в массиве только одна рифма, проверяем ее на полное совпадение с запросом
		//if i_dict+1 == j_dict && (len_word == len(new_dict[i_dict])) { // в массиве только одна рифма, проверяем ее на полное совпадение с запросом
		//fmt.Println(">>>единственная рифма равна запросу -", reverse(new_dict[i_dict]))
		return nil
	}

	if i_dict == j_dict {
		//fmt.Println(">>>idict is empty")
		return nil
	}

	// iindex := make(map[byte]int, 10)
	// d.index = iindex
	//fmt.Printf(">>>i_dict=%v, j_dict=%v, d.idx_pair.i=%v >>%v\n", i_dict, j_dict, d.idx_pair.i, d.dict[d.idx_pair.i:d.idx_pair.j])

	return &pair_int{i_dict + d.idx_pair.i, j_dict + d.idx_pair.i}
}

func (d *Dictionary2) Get_Ryth2(s string) string {
	sd := d.speed_search(reverse(s), s, true)
	if sd != "" {
		//fmt.Println("speed search", sd)
		return sd
	}
	temp_pair := new(pair_int)
	//idict := new(pair_int)
	ryth := ""
	length := len(s)
	//fmt.Println("\nRYTH:", s)
	for i := length - 1; i >= 0; i-- {
		count := length - i - 1
		////fmt.Printf("RYTH: idict:%v i:%v j:%v; count:%v\n", d.dict[d.idx_pair.i:d.idx_pair.j], d.idx_pair.i, d.idx_pair.j, count)
		if count == 9 && len(s) == 10 {
			//fmt.Println("RYTH: остался максимальный допустимый символ, завершаем поиск (исключаем полное совпадение)", s, d.dict[d.idx_pair.i:d.idx_pair.j])
			break
		}

		//исключаем из диапазона поиска строки с len < count (они в начале диапазона)
		new_in := 0
		for new_in = d.idx_pair.j - 1; new_in >= d.idx_pair.i; new_in-- {
			if len(d.dict[new_in]) > count {
				continue
			} else {
				//fmt.Println("RYTH:исключаем из диапазона поиска строки длиной больше count=", count)
				break
			}
		}
		if new_in == d.idx_pair.j-1 { // отсеяли весь массив
			break
		}
		if d.idx_pair.j-new_in-1 == 1 && s == reverse(d.dict[new_in+1]) {
			//fmt.Printf("RYTH: после удаления коротких рифм в массиве единственная рифма совпадает с запросом (исключаем), завершаем поиск. temp_pair.i=%v, temp_pair.j+%v %v//%v\n", new_in+1, d.idx_pair.j, s, reverse(d.dict[new_in+1]))
			break

		}
		if d.idx_pair.j-new_in-1 == 1 {
			//fmt.Printf("RYTH: после удаления коротких рифм в массиве единственная рифма НЕ совпадает с запросом, завершаем поиск. temp_pair.i=%v, temp_pair.j+%v %v//%v\n", new_in+1, d.idx_pair.j, s, reverse(d.dict[new_in+1]))
			d.idx_pair.i = new_in + 1
			break

		}

		d.idx_pair.i = new_in + 1

		//fmt.Printf("RYTH: idict: i:%v j:%v; count:%v\n", d.idx_pair.i, d.idx_pair.j, count)
		temp_pair = d.NewPair(s[i], count, length)
		if temp_pair == nil {
			//fmt.Println("RYTH: массиве нет рифм, завершаем поиск")
			break
		}
		//idict.i, idict.j = temp_pair.i, temp_pair.j

		if temp_pair.j-temp_pair.i == 1 && s == reverse(d.dict[temp_pair.i]) { //
			//fmt.Printf("RYTH: массиве единственная рифма совпадает с запросом (исключаем), завершаем поиск. temp_pair.i=%v, temp_pair.j+%v %v//%v\n", temp_pair.i, temp_pair.j, s, reverse(d.dict[temp_pair.i]))
			break
		}
		d.idx_pair.i, d.idx_pair.j = temp_pair.i, temp_pair.j
		temp_pair = nil
	}

	//fmt.Println("RYTH: out. string(idict.dict[0])/s/idict.dict: ", s, d.dict[d.idx_pair.i:d.idx_pair.j])
	//len_dict_temp := len(d.dict[d.idx_pair.i:d.idx_pair.j])
	for _, r := range d.dict[d.idx_pair.i:d.idx_pair.j] {
		if reverse(string(r)) != s { //
			// if len(s) > 8 { //&& i < len_dict_temp-1 {
			////fmt.Println("\nслово в массиве не совбадает с запросом - выводи", reverse(string(r)), s, d.dict[d.idx_pair.i:d.idx_pair.j])
			// 	//continue
			// }
			return reverse(string(r))
		}
	}
	for _, ds := range d.dict { //
		if ds[0] != s[length-1] {
			//fmt.Println("выводим любое слово НЕ рифму", reverse(string(ds)), ds[0] != s[length-1])
			return reverse(string(ds))
		}
	}

	return reverse(ryth)
}
func (d *Dictionary2) speed_search(sfull_revers, str string, isFirst bool) string {
	s := reverse(str)
	length := len(s)
	////fmt.Println(">>>", sfull_revers, s, isFirst)
	new_dict := d.dict[d.index[s[0]]:d.idx_pair.j]
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

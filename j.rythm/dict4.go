package main

import (
	"bufio"
	"io"
	"sort"
)

func run4(in io.Reader, out io.Writer) error {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	dict := NewDict(r)
	//full_index := dict.index
	words, err := get_arr(r)
	if err != nil {
		return err
	}

	for _, word := range *words {

		r := dict.Get_Ryth4(word, pair_int{0, len(dict.dict)})
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

func (d *Dictionary) NewPair4(ch byte, count int, len_word int, d_idx_pair pair_int) *pair_int {
	//fmt.Println("===NEW_DICT_BY_CHAR===\n get from dict: ", ch, "; by char: ", string(ch), "; count: ", count)
	// if len(d.dict) > 20 {
	// } else {
	//////// 	fmt.Println("===NEW_DICT_BY_CHAR===\n get from dict: ", d, "; by char: ", string(ch), "; count: ", count)
	// }

	isLast := len_word == count+1
	var i_dict, j_dict int

	//пропускаем короткие строки
	new_i := 0
	for new_i = d_idx_pair.i; new_i < d_idx_pair.j; new_i++ {
		if d.dict_sl[count][new_i] != 0 {
			break
		}
	}
	if new_i == d_idx_pair.j { //все строки короткие
		return nil
	}

	new_sl := d.dict_sl[count][new_i:d_idx_pair.j]

	i_dict = sort.Search(len(new_sl), func(i int) bool { return ch <= new_sl[i] })
	//fmt.Println("after search idict/new_i:", i_dict, new_i, "; new_sl:", new_sl) //, string(d.dict[i_dict]))
	i_dict = i_dict + d_idx_pair.i

	if d_idx_pair.j == i_dict { // символ не найден
		return nil
	}

	for j_dict = i_dict; j_dict < d_idx_pair.j; j_dict++ {
		//fmt.Println("idict", j_dict, ":", string(d.dict[j_dict]))
		//fmt.Println("есть рифма ? ch == d.dict[count][j]", ch, d.dict_sl[count][j_dict])
		if ch == d.dict_sl[count][j_dict] {
			//continue
		} else { //конец интервала в отсортированном массиве с нужным символом
			break
		}
	}

	for i := i_dict; i < j_dict; i++ {
		//fmt.Printf("совпадение рифмы и запроса? %v. слово в словаре: %v (%v)>> %v\n", (isLast && d.dict_sl[count+1][i] == 0), reverse(string(d.dict[i])), len(d.dict[i]), d.dict[i_dict:j_dict])
		if !(isLast && d.dict_sl[count+1][i] == 0) { //последний символ + длинны запроса и слова совпадают - рифму не учитываем (по условию)
			//fmt.Printf("совпадение рифмы и запроса. i=%v, i_dict=%v, j_dict=%v\n", i, i_dict, j_dict)
			//continue
			//} else { //конец интервала в отсортированном массиве с нужным символом
			i_dict = i
			break
		}

	}
	if i_dict+1 == j_dict && isLast && (d.dict_sl[count+1][i_dict] == 0) { // в массиве только одна рифма, проверяем ее на полное совпадение с запросом
		//if i_dict+1 == j_dict && (len_word == len(new_dict[i_dict])) { // в массиве только одна рифма, проверяем ее на полное совпадение с запросом
		//fmt.Println(">>>единственная рифма равна запросу -", reverse(d.dict[i_dict]))
		return nil
	}

	if i_dict == j_dict {
		//fmt.Println(">>>idict is empty")
		return nil
	}

	//fmt.Printf(">>>i_dict=%v, j_dict=%v, d_idx_pair.i=%v >>%v\n", i_dict, j_dict, d_idx_pair.i, d.dict[d_idx_pair.i:d_idx_pair.j])

	return &pair_int{i_dict, j_dict}
}

func (d *Dictionary) Get_Ryth4(str string, d_idx_pair pair_int) string {
	sd := d.speed_search(reverse(str), str, true, d_idx_pair)
	if sd != "" {
		//fmt.Println("speed search", sd)
		return sd
	}
	temp_pair := new(pair_int)
	ryth := ""
	length := len(str)
	//fmt.Println("\nRYTH:", str)
	for i, s := range reverse(str) { //length - 1; i >= 0; i-- {
		count := i
		////fmt.Printf("RYTH: idict:%v i:%v j:%v; count:%v\n", d.dict[d_idx_pair.i:d_idx_pair.j], d_idx_pair.i, d_idx_pair.j, count)
		if count == 9 && len(str) == 10 {
			//fmt.Println("RYTH: остался максимальный допустимый символ, завершаем поиск (исключаем полное совпадение)", s, d.dict[d_idx_pair.i:d_idx_pair.j])
			break
		}

		//исключаем из диапазона поиска строки с len < count (они в начале диапазона)
		new_in := 0
		for new_in = d_idx_pair.j - 1; new_in >= d_idx_pair.i; new_in-- {
			if len(d.dict[new_in]) > count {
				continue
			} else {
				//fmt.Println("RYTH:исключаем из диапазона поиска строки длиной больше count=", count)
				break
			}
		}
		if new_in == d_idx_pair.j-1 { // отсеяли весь массив
			break
		}
		if d_idx_pair.j-new_in-1 == 1 && str == reverse(d.dict[new_in+1]) {
			//fmt.Printf("RYTH: после удаления коротких рифм в массиве единственная рифма совпадает с запросом (исключаем), завершаем поиск. temp_pair.i=%v, temp_pair.j+%v %v//%v\n", new_in+1, d_idx_pair.j, s, reverse(d.dict[new_in+1]))
			break

		}
		if d_idx_pair.j-new_in-1 == 1 {
			//fmt.Printf("RYTH: после удаления коротких рифм в массиве единственная рифма НЕ совпадает с запросом, завершаем поиск. temp_pair.i=%v, temp_pair.j+%v %v//%v\n", new_in+1, d_idx_pair.j, s, reverse(d.dict[new_in+1]))
			d_idx_pair.i = new_in + 1
			break

		}

		d_idx_pair.i = new_in + 1

		//fmt.Printf("RYTH: idict: i:%v j:%v; count:%v\n", d_idx_pair.i, d_idx_pair.j, count)
		temp_pair = d.NewPair4(byte(s), count, length, d_idx_pair)
		if temp_pair == nil {
			//fmt.Println("RYTH: массиве нет рифм, завершаем поиск")
			break
		}
		//idict.i, idict.j = temp_pair.i, temp_pair.j

		if temp_pair.j-temp_pair.i == 1 && str == reverse(d.dict[temp_pair.i]) { //
			//fmt.Printf("RYTH: массиве единственная рифма совпадает с запросом (исключаем), завершаем поиск. temp_pair.i=%v, temp_pair.j+%v %v//%v\n", temp_pair.i, temp_pair.j, s, reverse(d.dict[temp_pair.i]))
			break
		}
		d_idx_pair.i, d_idx_pair.j = temp_pair.i, temp_pair.j
		temp_pair = nil
	}

	//fmt.Println("RYTH: out. string(idict.dict[0])/s/idict.dict: ", str, d.dict[d_idx_pair.i:d_idx_pair.j])
	//len_dict_temp := len(d.dict[d_idx_pair.i:d_idx_pair.j])
	for _, r := range d.dict[d_idx_pair.i:d_idx_pair.j] {
		if reverse(string(r)) != str { // вывод результата не совпадающего с запросом
			return reverse(string(r))
		}
	}
	for _, ds := range d.dict { //
		if ds[0] != str[length-1] {
			//fmt.Println("выводим любое слово НЕ рифму", reverse(string(ds)), ds[0] != str[length-1])
			return reverse(string(ds))
		}
	}

	return reverse(ryth)
}

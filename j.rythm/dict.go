package main

func NewDictByChar(d *Dictionary2, ch byte, count int, len_word int) *Dictionary2 {
	//fmt.Println("===NEW_DICT_BY_CHAR===\n get from dict: ", d, "; by char: ", string(ch), "; count: ", count)
	isLast := len_word == count+1
	t_slice := []string{}
	for _, r := range d.dict { //перед след поиском убираем из массива варианты с длинной менее count
		if len(r) > count {
			t_slice = append(t_slice, r)
		}
	}
	if len(t_slice) == 0 { //
		//fmt.Println("t_lice empty")
		return nil
	} else if len(t_slice) == 1 { //
		//fmt.Println("len t_lice == 1")
		if isLast && (ch == byte(t_slice[0][count])) { //полное совпадение (исключаем)
			return nil
		}
		d.dict = t_slice
		return d
	}
	d.dict = t_slice

	var idict []string
	if count > 0 { //не первоначальный массив (уже обработанный раньше)
		//строим индексы по заданному символу
		for idx, c := range d.dict {
			if _, inMap := d.index[c[count]]; inMap {
				continue
			}
			d.index[c[count]] = idx
		}
	}
	i_dict, inMap := d.index[ch]
	if !inMap {
		//fmt.Println("not in map", d, reverse(string(d.dict[j])), string(ch), ch, count, j)
		return nil
	}
	for j := i_dict; j < len(d.dict); j++ {
		//fmt.Println("idict", j, ":", reverse(string(d.dict[j])))
		len_dictj := len(d.dict[j])
		//if len_dictj > count {
		//fmt.Println("ch == d.dict[j][count]", ch, d.dict[j][count])
		if ch == d.dict[j][count] {
			//fmt.Println("совпадение рифмы и запроса? ", (isLast && (len_word == len_dictj)), isLast, len_word, len_dictj, reverse(string(d.dict[j])))
			if !(isLast && (len_word == len_dictj)) { //последний символ + длинны запроса и слова совпадают - рифму не учитываем (по условию)
				idict = append(idict, d.dict[j])
			}
		} else { //конец интервала в отсортированном массиве с нужным символом
			break
		}
		//}
	}

	if len(idict) == 0 {
		////		fmt.Println(">>>idict is empty")
		return nil
	}
	iindex := make(map[byte]int)
	//fmt.Println(">>>", idict, iindex)

	return &Dictionary2{idict, iindex, pair_int{0, 0}}
}

func (d *Dictionary2) Get_Ryth(s string) string {
	var temp_pair *Dictionary2
	idict := d
	ryth := ""
	count := 0
	length := len(s)
	//fmt.Println("RYTH:", s)
	for i := length - 1; i >= 0; i-- {
		//fmt.Println("RYTH: idict", idict, " i:", i)
		temp_pair = NewDictByChar(idict, s[i], count, length)
		if temp_pair == nil {
			//fmt.Println("RYTH: массиве нет рифм, завершаем поиск")
			break
		}
		idict = temp_pair
		// if len(idict.dict) == 1 { //
		//// 	fmt.Println("RYTH: массиве только одна рифма, завершаем поиск")
		// 	break
		// }
		if count == 9 {
			//fmt.Println("RYTH: остался максимальный допустимый символ, завершаем поиск (исключаем полное совпадение)")
			break
		}
		temp_pair = nil
		count++
	}

	//fmt.Println("RYTH: out. string(idict.dict[0])/s/idict.dict: ", string(idict.dict[0]), s, idict.dict)
	for _, r := range idict.dict {
		if reverse(string(r)) != s { //
			//fmt.Println("слово в массиве не совбадает с запросом - выводи", reverse(string(r)), s)
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

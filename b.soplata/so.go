package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fin, _ := os.Open("./tests/06")
	err := so(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}

}
func so(in io.Reader, out io.Writer) error {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	str_t, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	t, err := strconv.Atoi(strings.Trim(str_t, "\r\n"))
	if err != nil {
		return fmt.Errorf("first line is not int.\n %v", err)
	}
	for i := 0; i < t; i++ {
		str_f, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		f, _ := strconv.Atoi(strings.Trim(str_f, "\r\n"))
		//log.Println(t, f)
		str_af, err := r.ReadString('\n')
		if err != nil {
			return err
		}

		af := strings.Split(strings.Trim(str_af, "\r\n"), " ")

		if len(af) != f {
			return fmt.Errorf("expectes %d numbers: %v", f, af)
		}
		gds := make(map[int]int)
		for _, s := range af {
			ns, _ := strconv.Atoi(s)
			gds[ns]++
		}
		sum := 0
		for k, v := range gds {
			free := v / 3
			sum += k * (v - free)
		}

		w.WriteString(fmt.Sprintf("%d\n", sum))
	}
	return nil
}

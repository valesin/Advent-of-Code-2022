package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	calorie := sommaBlocchi()
	tot := sommaPrimiTre(calorie)
	fmt.Println(tot)
}
func sommaBlocchi() (listaSomme []int) {
	var totBlocco int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "%" {
			break
		}
		linea, _ := strconv.Atoi(s.Text())
		totBlocco += linea
		if s.Text() == "" {
			listaSomme = append(listaSomme, totBlocco)
			totBlocco = 0
		}
	}
	return
}
func sommaPrimiTre(listaSomme []int) (tot int) {
		var max, max2, max3 int
	for _, v := range listaSomme {
		if v > max {
			max3 = max2
			max2 = max
			max = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}
	return max + max2 + max3
}

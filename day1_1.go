package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//var sommaCal int
	tot := sommaBlocchi()
	fmt.Println(tot)
}
func sommaBlocchi() (maxBlocchi int) {
	var totBlocco int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "%" {
			break
		}
		linea, _ := strconv.Atoi(s.Text())
		totBlocco += linea
		if s.Text() == "" {

			if totBlocco > maxBlocchi {
				maxBlocchi = totBlocco
			}
			totBlocco = 0
		}
	}

	return

}

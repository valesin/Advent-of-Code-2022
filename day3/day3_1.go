package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var totPriorità int
	righe := leggiRighe()
	for _, riga := range righe {
		totPriorità += calcolaPriorità(trovaRipetizione(riga))
	}
	fmt.Println(totPriorità)
}
func leggiRighe() (righe []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		righe = append(righe, scanner.Text())
	}
	return
}
func trovaRipetizione(riga string) (rip byte) {
	for i := 0; i < len(riga)/2; i++ {
		for j := len(riga) / 2; j < len(riga); j++ {
			if riga[i] == riga[j] {
				return riga[i]
			}
		}
	}
	return
}
func calcolaPriorità(rip byte) (priorità int) {
	if unicode.IsUpper(rune(rip)) {
		priorità = int(rip) - 38
	} else {
		priorità = int(rip) - 96
	}
	return

}

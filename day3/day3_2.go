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
	byteComuni := trovaByteComune3(righe)
	for _, e := range byteComuni {
		totPriorità += calcolaPriorità(e)
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
func trovaByteComune3(righe []string) (byteComuni []byte) {
	for i := 0; i < len(righe); i += 3 {
	PrimaRiga:
		for j := 0; j < len(righe[i]); j++ {
			for q := 0; q < len(righe[i+1]); q++ {
				if righe[i][j] == righe[i+1][q] {
					for k := 0; k < len(righe[i+2]); k++ {
						if righe[i+1][q] == righe[i+2][k] {
							byteComuni = append(byteComuni, righe[i][j])
							break PrimaRiga
						}
					}
				}
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

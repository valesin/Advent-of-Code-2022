package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var tot int
	for _, v := range leggiTesto() {
		tot += calcolaPunteggioRiga(v)
	}
	fmt.Println(tot)
}

func leggiTesto() (righe []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		righe = append(righe, scanner.Text())
	}
	return
}
func calcolaPunteggioRiga(riga string) (punteggio int) {
	mossaAvversario, mossaMia := riga[0], riga[2]-23
	if mossaMia == 'A' {
		punteggio += 1
		if mossaAvversario == 'A' {
			punteggio += 3
		} else if mossaAvversario == 'B' {
		} else if mossaAvversario == 'C' {
			punteggio += 6
		}
	} else if mossaMia == 'B' {
		punteggio += 2
		if mossaAvversario == 'A' {
			punteggio += 6
		} else if mossaAvversario == 'B' {
			punteggio += 3
		} else if mossaAvversario == 'C' {
		}
	} else if mossaMia == 'C' {
		punteggio += 3
		if mossaAvversario == 'A' {
		} else if mossaAvversario == 'B' {
			punteggio += 6
		} else if mossaAvversario == 'C' {
			punteggio += 3
		}
	}
	return
}

// b vince su a, c vince su b, a vince su c.

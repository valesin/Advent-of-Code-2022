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
	mossaAvversario, esito := riga[0], riga[2]
	if esito == 'Y' {
		punteggio += 3
		if mossaAvversario == 'A' {
			punteggio += 1
		} else if mossaAvversario == 'B' {
			punteggio += 2
		} else if mossaAvversario == 'C' {
			punteggio += 3
		}
	} else if esito == 'X' {
		if mossaAvversario == 'A' {
			punteggio += 3
		} else if mossaAvversario == 'B' {
			punteggio += 1
		} else if mossaAvversario == 'C' {
			punteggio += 2
		}
	} else if esito == 'Z' {
		punteggio += 6
		if mossaAvversario == 'A' {
			punteggio += 2
		} else if mossaAvversario == 'B' {
			punteggio += 3
		} else if mossaAvversario == 'C' {
			punteggio += 1
		}
	}
	return
}

// b vince su a, c vince su b, a vince su c.

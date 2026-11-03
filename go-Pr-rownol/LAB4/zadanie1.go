package main

import (
	"fmt"
	"time"
)

// Funkcja sprawdzająca czy liczba jest pierwszą (najprostsza, wolna metoda)
func czyPierwsza(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func szukajPierwszej(kanalKonca chan bool) {
	fmt.Println("Start szukania liczby pierwszej...")
	kandydat := 2
	ostatniaPierwsza := 2

	for {
		// Sprawdzamy, czy w kanale końca pojawił się sygnał (bez blokowania dzięki default)
		select {
		case <-kanalKonca:
			fmt.Println("Otrzymano sygnał końca. Przerywam obliczenia.")
			fmt.Printf("Ostatnia wyliczona liczba pierwsza: %d\n", ostatniaPierwsza)
			return
		default:
			// Kontynuujemy obliczenia
		}

		if czyPierwsza(kandydat) {
			ostatniaPierwsza = kandydat
			time.Sleep(1 * time.Millisecond) // Krótkie uśpienie po znalezieniu
		}
		kandydat++
	}
}

func main() {
	kanalCzasNaKoniec := make(chan bool)

	go szukajPierwszej(kanalCzasNaKoniec)

	// Czekamy ok. 10ms
	time.Sleep(10 * time.Millisecond)

	// Wysyłamy sygnał końca
	kanalCzasNaKoniec <- true

	// Czekamy chwilę, aby wątek zdążył wypisać komunikat przed zakończeniem programu
	time.Sleep(1 * time.Second)
	fmt.Println("Koniec programu main().")
}

package main

import (
	"context"
	"fmt"
	"time"
)

func czyPierwsza2(n int) bool {
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

func szukajPierwszejZContext(ctx context.Context) {
	fmt.Println("Start szukania liczby pierwszej z kontekstem...")
	kandydat := 2
	ostatniaPierwsza := 2

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Kontekst zgłasza koniec. Przerywam obliczenia.")
			fmt.Printf("Ostatnia wyliczona liczba pierwsza: %d\n", ostatniaPierwsza)
			return
		default:
			// brak blokowania, idziemy dalej
		}

		if czyPierwsza2(kandydat) {
			ostatniaPierwsza = kandydat
			time.Sleep(1 * time.Millisecond)
		}
		kandydat++
	}
}

func main() {
	// Tworzymy kontekst z automatycznym timeoutem po 10ms
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	go szukajPierwszejZContext(ctx)

	// Czekamy dłużej niż 10ms, aby wątek zdążył zareagować i zakończyć działanie
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Koniec programu main().")
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func czyPierwsza3(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Funkcja-wątek szukająca najmniejszej liczby pierwszej > prog
func worker(ctx context.Context, cancel context.CancelFunc, id int, startVal int, prog int, wg *sync.WaitGroup) {
	defer wg.Done()

	kandydat := startVal
	ostatniaZnaleziona := 0

	for {
		select {
		case <-ctx.Done():
			// Kontekst został zamknięty przez kogoś szybszego
			return
		default:
		}

		if kandydat > prog && czyPierwsza3(kandydat) {
			ostatniaZnaleziona = kandydat
			// Znalazł! Informujemy inne wątki poprzez anulowanie kontekstu
			fmt.Printf("Wątek %d znalazł szukaną liczbę pierwszą: %d\n", id, ostatniaZnaleziona)
			cancel()
			return
		}

		kandydat++
		// Delikatne spowolnienie, aby zaobserwować współbieżność
		time.Sleep(10 * time.Microsecond)
	}
}

func main() {
	prog := 1000
	liczbaWatkow := 4

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Uruchamiamy kilka wątków z różnymi punktami startowymi (lub tym samym, zależnie od koncepcji wyścigu)
	for i := 1; i <= liczbaWatkow; i++ {
		wg.Add(1)
		// Każdy wątek może startować np. od innej bazy albo wszystkie od 2
		go worker(ctx, cancel, i, 2, prog, &wg)
	}

	wg.Wait()
	fmt.Println("Wszystkie wątki zakończyły działanie.")
}

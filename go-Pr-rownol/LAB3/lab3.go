package main

import (
	"fmt"
)

// 2. Funkcja wykonywana w osobnych gorutynach
func sumator(id int, tablica []float64, kanalPartii chan int, kanalSum chan float64) {
	dlugoscCzesci := len(tablica) / 4

	for nrCzesci := range kanalPartii {
		start := nrCzesci * dlugoscCzesci
		end := start + dlugoscCzesci

		// Jeśli to ostatnia część, uwzględniamy ewentualny resztkowy rozmiar
		if nrCzesci == 3 {
			end = len(tablica)
		}

		sumaCzesciowa := 0.0
		for i := start; i < end; i++ {
			sumaCzesciowa += tablica[i]
		}

		fmt.Printf("[Worker %d] Zsumowano czesc %d (indeksy %d..%d), suma: %.2f\n",
			id, nrCzesci, start, end-1, sumaCzesciowa)

		// 3. Przekazanie sumy do drugiego kanału
		kanalSum <- sumaCzesciowa
	}
}

func main() {
	// 1. Tworzenie tablicy 10 000 elementów i wypełnianie +1/3 oraz -1/3
	rozmiar := 10000
	tablica := make([]float64, rozmiar)

	polowa := rozmiar / 2
	for i := 0; i < polowa; i++ {
		tablica[i] = 1.0 / 3.0
	}
	for i := polowa; i < rozmiar; i++ {
		tablica[i] = -1.0 / 3.0
	}

	// 2. Tworzenie kanałów o długości (buforze) 4 oraz 2 workerów
	kanalPartii := make(chan int, 4)
	kanalSum := make(chan float64, 4)

	// Uruchomienie 2 wątków-funkcji (workerów)
	go sumator(1, tablica, kanalPartii, kanalSum)
	go sumator(2, tablica, kanalPartii, kanalSum)

	// Wysyłanie 4 numerów części do kanału (0, 1, 2, 3)
	for i := 0; i < 4; i++ {
		kanalPartii <- i
	}
	// Zamknięcie kanału partii sygnalizuje workerom koniec nowych zadań
	close(kanalPartii)

	// 3. i 4. Odbieranie sum cząstkowych z drugiego kanału oraz wyliczanie sumy ostatecznej
	sumaOstateczna := 0.0
	for i := 0; i < 4; i++ {
		czastkowa := <-kanalSum
		fmt.Printf("[main] Otrzymano sume czastkowa: %.2f\n", czastkowa)
		sumaOstateczna += czastkowa
	}

	// Wyświetlenie ostatecznego wyniku
	fmt.Printf("\nOstateczna suma calej tablicy: %.2f\n", sumaOstateczna)
}

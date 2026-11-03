package main

import (
	"fmt"
	"sync"
)

// Globalna tablica oraz mechanizmy synchronizacji
var tab [10000]float64
var mut sync.RWMutex
var wg sync.WaitGroup

// Zadanie 2: Funkcja statystyka()
// Wyświetla sumę oraz maksymalną wartość, iterując od końca do początku.
func statystyka() {
	suma := 0.0
	// Ustawiamy maxVal na pierwszy element od końca
	maxVal := tab[len(tab)-1]

	for i := len(tab) - 1; i >= 0; i-- {
		suma += tab[i]
		if tab[i] > maxVal {
			maxVal = tab[i]
		}
	}
	fmt.Printf("[STATYSTYKA] Suma: %.2f | Maksymalna wartość: %.2f\n", suma, maxVal)
}

// Zadanie 3: Wątko-funkcja czytacz()
// Blokuje odczyt za pomocą RLock/RUnlock i uruchamia statystyka()
func czytacz(id int) {
	defer wg.Done()
	fmt.Printf("Czytacz %d: czeka na RLock...\n", id)
	mut.RLock()
	fmt.Printf("Czytacz %d: uzyskany RLock, uruchamiam statystykę.\n", id)

	statystyka()

	mut.RUnlock()
	fmt.Printf("Czytacz %d: zwolniono RUnlock.\n", id)
}

// Zadanie 4 i 5: Wątko-funkcja pisarz-mnoznik()
// Mnoży wszystkie elementy tablicy przez 2, blokując dostęp za pomocą Lock/Unlock
func pisarzMnoznik(id int) {
	defer wg.Done()
	fmt.Printf("Pisarz %d: czeka na pełną blokadę (Lock)...\n", id)
	mut.Lock()
	fmt.Printf("Pisarz %d: uzyskany Lock, mnożę elementy tablicy przez 2...\n", id)

	for i := 0; i < len(tab); i++ {
		tab[i] *= 2.0
	}

	mut.Unlock()
	fmt.Printf("Pisarz %d: zwolniono Lock.\n", id)
}

func main() {
	// --- ZADANIE 1 ---
	// Wypełnienie tablicy w połowie +1 oraz w połowie -1
	polowa := len(tab) / 2
	for i := 0; i < len(tab); i++ {
		if i < polowa {
			tab[i] = 1.0
		} else {
			tab[i] = -1.0
		}
	}

	fmt.Println("Stan początkowy tablicy (przed uruchomieniem wątków):")
	statystyka()
	fmt.Println("--------------------------------------------------")

	// --- ZADANIE 5 ---
	// Jednoczesne uruchomienie 3 czytaczy i 2 pisarzy-mnożników
	numCzytaczy := 3
	numPisarzy := 2

	// Dodajemy do WaitGroup łączną liczbę wątków (3 + 2 = 5)
	wg.Add(numCzytaczy + numPisarzy)

	// Uruchomienie 3 wątków czytaczy
	for i := 1; i <= numCzytaczy; i++ {
		go czytacz(i)
	}

	// Uruchomienie 2 wątków pisarzy
	for i := 1; i <= numPisarzy; i++ {
		go pisarzMnoznik(i)
	}

	// Oczekiwanie na zakończenie wszystkich wątków
	wg.Wait()

	fmt.Println("--------------------------------------------------")
	fmt.Println("Wszystkie wątki zakończyły pracę. Ostateczna statystyka:")
	statystyka()
}

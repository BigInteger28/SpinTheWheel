package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	file, err := os.Open("keuzes.txt") // Vervang "keuzes.txt" met het pad naar je bestand
	if err != nil {
		fmt.Println("Fout bij het openen van het bestand:", err)
		return
	}
	defer file.Close()

	var keuzes []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keuzes = append(keuzes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Fout bij het lezen van het bestand:", err)
		return
	}

	// Kies een willekeurige prijs
	if len(keuzes) > 0 {
		for {
			fmt.Print("\nLet's go girl: ")
			fmt.Scanln()
			spin := rand.Intn(len(keuzes))
			// Print de gekozen prijs
			fmt.Println("\n",keuzes[spin])
		}
	} else {
		fmt.Println("Geen keuzes gevonden in het bestand.")
	}


}

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
	var bestand string = "keuzes"
	for {		
		var keuze string
		fmt.Println("\n1. BESTAND KIEZEN")
		fmt.Println("SPIN THE WHEEL")
		fmt.Print(": ")
		fmt.Scanln(&keuze)
		if keuze == "1" {
			fmt.Print("Ik kies bestand: ")
			fmt.Scanln(&bestand)
		} else {			
			file, err := os.Open(bestand+".txt")
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
			if len(keuzes) > 0 {
			spin := rand.Intn(len(keuzes))
			fmt.Println("\n",keuzes[spin])
			} else {
				fmt.Println("Geen keuzes gevonden in het bestand.")
			}			
		}
	}
}

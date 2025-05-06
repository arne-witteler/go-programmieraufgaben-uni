// Schreiben Sie ein Programm, das eine Uhrzeit (Stunde 0 bis 23, Minute 0 bis 59)
// einliest und den Winkel des Stunden- und Minutenzeigers zu dieser Uhrzeit ausgibt.
// Die Winkel sind 0°, wenn die Zeiger oben stehen, also um 0 und 12 Uhr.

package main

import "fmt"

func main() {

	var stunde int
	var minute int
	var stundenWinkel float64
	var minutenWinkel float64

	fmt.Println("Geben Sie bitte eine Uhrzeit an, indem Sie zunächst die Stunde (von 0 bis 23) und dann die Minute (von 0 bis 59) eingeben:")
		fmt.Scan(&stunde)
		fmt.Scan(&minute)

	fmt.Println()
	
	stundenWinkel = float64(360) / 24 * float64(stunde)
	minutenWinkel = float64(360) / 60 * float64(minute)
	
	fmt.Printf("Zeigerstellung um %02d:%02d Uhr:\n", stunde, minute)
	fmt.Printf("Winkel des Stundenzeigers: %.2f°\n", stundenWinkel)
	fmt.Printf("Winkel des Minutenzeigers: %.2f°\n", minutenWinkel)
}
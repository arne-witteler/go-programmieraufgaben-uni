// Schreiben Sie ein Programm, daß die Längen der drei Seiten eines Quaders einliest
// und folgende Größen dieses Quaders ausgibt:
// • Volumen
// • Kantensumme
// • Oberfläche
// • Umkugelradius
// • Raumdiagonale

package main

import (
	"fmt"
	"math"
)

func main() {

	var laenge float64
	var breite float64
	var hoehe float64
	var volumen float64
	var kantensumme float64
	var oberflaeche float64
	var umkugelradius float64
	var raumdiagonale float64
	
	fmt.Println("Bitte geben Sie die drei Seitenlängen des Quaders ein: ")
		fmt.Scan(&laenge)
		fmt.Scan(&breite)
		fmt.Scan(&hoehe)

		volumen = laenge * breite * hoehe
		kantensumme = 4 * (laenge + breite + hoehe)
		oberflaeche = 2 * (laenge * breite + laenge * hoehe + breite * hoehe)
		raumdiagonale = math.Sqrt(math.Pow(laenge, 2) + math.Pow(breite, 2) + math.Pow(hoehe, 2))
		umkugelradius = (raumdiagonale) / 2 

		fmt.Println("Ein", laenge, "×", breite, "×", hoehe, "Quader hat die geometrischen Größen: ")
		fmt.Println("Volumen:", volumen)
		fmt.Println("Kantensumme:", kantensumme)
		fmt.Println("Oberfläche:", oberflaeche)
		fmt.Printf("Umkugelradius: %.2f\n", umkugelradius)
		fmt.Printf("Raumdiagonale: %.2f\n", raumdiagonale)
}
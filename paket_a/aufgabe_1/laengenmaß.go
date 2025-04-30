// Schreibe ein Programm, das eine Länge in Metern einliest und diese umrechnet in:
// • mm
// • km
// • Zoll (1 Zoll = 2.54 cm)
// • Seemeilen (1 sm = 1852 m)
// • Lichtjahre (1 Lj = 9 460 730 472 580 800 m)

package main

import "fmt"

func main() {

	var meter int
	var mm float64
	var km float64
	var zoll float64
	var sm float64
	var lj float64

	fmt.Println("Gib eine Länge (in Metern) ein: ")
		fmt.Scan(&meter)

	mm = float64(meter) * 1000
	km = float64(meter) / 1000
	zoll = float64(meter) * 39.37
	sm = float64(meter) / 1852
	lj = float64(meter) / 9.461e+15

	fmt.Println(meter, "Meter entsprechen:")
	fmt.Println(mm, "mm")
	fmt.Println(km, "km")
	fmt.Println(zoll, "Zoll")
	fmt.Println(sm, "sm")
	fmt.Println(lj, "Lj")
}
// Schreibe ein Programm, das eine Länge in Metern einliest und diese umrechnet in:
// • mm
// • km
// • Zoll (1 Zoll = 2.54 cm)
// • Seemeilen (1 sm = 1852 m)
// • Lichtjahre (1 Lj = 9 460 730 472 580 800 m)

package main

import "fmt"

func main() {

	const zollProMeter = 39.37007874
	const meterProSeemeile = 1852.0
	const meterProLichtjahr = 9.4607304725808e15

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
	zoll = float64(meter) * zollProMeter
	sm = float64(meter) / meterProSeemeile
	lj = float64(meter) / meterProLichtjahr

	fmt.Println(meter, "Meter entsprechen:")
	fmt.Println(mm, "mm")
	fmt.Println(km, "km")
	fmt.Println(zoll, "Zoll")
	fmt.Println(sm, "sm")
	fmt.Println(lj, "Lj")
}
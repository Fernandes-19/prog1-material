package aufgabe3

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.
func CountSquares(list []int) int {

	if len(list) == 0 {
		return 0
	}

	for x := 0; x <= list[0]; x++ {
		if x*x == list[0] {
			return 1 + CountSquares(list[1:])
		}

	}

	return CountSquares(list[1:])
}

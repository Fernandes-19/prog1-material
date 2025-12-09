package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.
func DuplicateSinglets(list []int) []int {

	var result []int

	for _, v := range list {
		count := 0
		for _, x := range list {
			if x == v {
				count++
			}
		}
		if count == 1 {
			result = append(result, v, v)
		} else {
			result = append(result, v)
		}
	}
	return result
}

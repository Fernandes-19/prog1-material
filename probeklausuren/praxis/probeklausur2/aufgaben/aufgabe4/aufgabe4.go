package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	var res []int

	// Auffüll-Werte (letztes Element oder 0)
	last1, last2 := 0, 0
	if len(l1) > 0 {
		last1 = l1[len(l1)-1]
	}
	if len(l2) > 0 {
		last2 = l2[len(l2)-1]
	}

	// so viele Durchläufe wie die längere Liste
	for i := 0; i < len(l1) || i < len(l2); i++ {
		a := last1
		if i < len(l1) {
			a = l1[i]
		}
		b := last2
		if i < len(l2) {
			b = l2[i]
		}
		res = append(res, a+b)
	}
	return res
}

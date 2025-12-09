package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	result := []int{}
	for i := 0; i < len(l1); i++ {
		if Contains(l1[i], l2) {
			result = append(result, l1[i])

		}

	}
	for i := 0; i < len(l2); i++ {
		if Contains(l2[i], l1) {
			result = append(result, l2[i])
		}
	}
	return result
}
func Contains(a int, list []int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == a {
			return false
		}
	}
	return true
}

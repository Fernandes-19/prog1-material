package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// FilterDigits erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Ziffern entfernt werden.
// Alle anderen Zeichen sollen unverändert bleiben.
func FilterDigits(s string) string {
	//a,b,c,d,e,f,g,h,i,j:="0","1","2","3","4","5","6","7","8","9"

	if s == "" {
		return ""
	}

	if s[0] < '0' || s[0] > '9' {
		return string(s[0]) + FilterDigits(s[1:])
	}

	return FilterDigits(s[1:])
}

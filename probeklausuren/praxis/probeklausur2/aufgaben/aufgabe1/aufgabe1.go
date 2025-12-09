package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// PrefixBelow10 erwartet eine Liste "list" von Zahlen und liefert
// die längste Teil-Liste, mit der "list" beginnt und die nur Zahlen < 10 enthält.
func PrefixBelow10(list []int) []int {
	//a,b,c,d,e,f,g,h,i,j:=0,1,2,3,4,5,6,7,8,9

	if len(list) == 0 {
		return []int{}

	}

	if list[0] > 9 {
		return []int{}
	}
	if list[0] < 10 {
		return append([]int{list[0]}, PrefixBelow10(list[1:])...)
	}
	return PrefixBelow10(list[1:])
}

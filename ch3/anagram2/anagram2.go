// gopl Exercise 3.12
// Avery VanKirk
package anagram2


//anagram using rune map and range iterator
func isAnagram(s1 string, s2 string) bool {

	lc := make(map[rune]int)
	for _,letter := range s1 {
		lc[letter]++
	}
	for _,letter := range s2 {
		lc[letter]--
	}
	for _, n := range lc {
		if n != 0 {
			return false			
		}
	}
	return true
}

//anagram using byte map and for loop
func isAnagramOld(s1 string, s2 string) bool {

	lc := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		lc[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		lc[s2[i]]--
	}

	for _, n := range lc {
		if n != 0 {
			fmt.Println("No Anagram")
			return false			
		}
	}
	fmt.Println("Is Anagram")
	return true
}

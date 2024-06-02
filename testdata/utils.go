package testdata

func Sum(a, b int) int {
	return a + b
}

func ReverseWord(word string) string {
	var reversed string
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}
	return reversed
}

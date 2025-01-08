package iteration

func Repeat(word string, count int) string {
	var repeat string
	for i := 0; i < count; i++ {
		repeat += word
	}
	return repeat
}

package iteration

func Repeat(char string, time int) string {
	var repeated string
	for i := 0; i < time; i++ {
		repeated += char
	}
	return repeated
}

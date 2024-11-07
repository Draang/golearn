package ciclos

func Repeat(character string, num int) string {
	var reap string
	for i := 0; i < num; i++ {
		reap += character
	}
	return reap
}

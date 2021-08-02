package iteration

func Repeat(s string, repeats int) string {
	var repeated string
	for i := 0; i < repeats; i++ {
		repeated = repeated + s
	}

	return repeated
}

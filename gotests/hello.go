package gotests

func Hello(s string) string {
	if s == "" {
		return "Hello Dude"
	}
	return "Hello " + s
}

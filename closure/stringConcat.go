package main

func StringConcat() func(string) string {
	text := ""
	return func(s string) string {
		text += s + " "
		return text
	}
}

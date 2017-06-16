package main

import "fmt"

func main() {
	var regex, text string
	regex = ".*e+c$"
	text = "abeec"
	fmt.Println(regex)
	fmt.Println(text)
	fmt.Println(match(regex, text))
}

func match(regex, text string) bool {
	if len(regex) == 0 || len(text) == 0 {
		return false
	}

	if regex[0] == '^' {
		return matchhere(regex[1:], text)
	}

	if matchhere(regex, text) {
		return true
	}

	for len(text) != 0 {
		text = text[1:]
		if matchhere(regex, text) {
			return true
		}
	}

	return false
}

func matchhere(regex, text string) bool {
	if len(regex) == 0 {
		return true
	}

	if len(regex) == 1 && regex[0] == '$' {
		return len(text) == 0
	}

	if len(regex) >= 2 && regex[1] == '*' {
		return matchstar(regex[0], regex[2:], text)
	}

	if len(regex) >= 2 && regex[1] == '+' {
		return matchplus(regex[0], regex[2:], text)
	}

	if len(text) != 0 && (regex[0] == '.' || regex[0] == text[0]) {
		return matchhere(regex[1:], text[1:])
	}

	return false
}

func matchstar(c byte, regex, text string) bool {
	if matchhere(regex, text) {
		return true
	}

	for len(text) != 0 && (c == text[0] || c == '.') {
		text = text[1:]
		if matchhere(regex, text) {
			return true
		}
	}

	return false
}

func matchplus(c byte, regex, text string) bool {
	for len(text) != 0 && (c == text[0] || c == '.') {
		text = text[1:]
		if matchhere(regex, text) {
			return true
		}
	}

	return false
}

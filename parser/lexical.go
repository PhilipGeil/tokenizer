package parser

import "strings"

var keywords = []string{"class", "constructor", "function", "method", "field", "static", "var", "int", "char", "boolean", "void", "true", "false", "null", "this", "let", "do", "if", "else", "while", "return"}

func startsWithKeyword(line string) string {
	for _, v := range keywords {
		if strings.HasPrefix(line, v) {
			return v
		}
	}
	return ""
}

var symbols = []string{"{", "}", "(", ")", "[", "]", ".", ",", ";", "+", "-", "*", "/", "&", "|", "<", ">", "=", "~"}

func containsSymbol(s string) bool {
	for _, v := range symbols {
		if v == s {
			return true
		}
	}
	return false
}

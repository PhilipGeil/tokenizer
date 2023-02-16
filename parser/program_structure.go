package parser

import (
	"strings"
)

var classVarDec = []string{"static", "field"}
var subroutineDec = []string{"constructor", "function", "method"}
var statement = []string{"let", "if", "while", "do", "return"}
var op = []string{"+", "-", "*", "/", "&amp;", "|", "&lt;", "&gt;", "="}
var unaryOp = []string{"-", "~"}

func isClassVarDec(line string) bool {
	for _, v := range classVarDec {
		if strings.Contains(line, v) {
			return true
		}
	}
	return false
}

func ParseClassVarDec(lines []string, index int) (res []string, idx int) {
	// append lines
	res = append(res, "  <classVarDec>")
	for i := index; i < len(lines); i++ {
		if strings.Contains(lines[i], ";") {
			res = append(res, "	"+lines[i])
			idx = i
			break
		}
		res = append(res, "	"+lines[i])
	}
	res = append(res, "  </classVarDec>")
	return
}

func isSubroutineDec(line string) bool {
	for _, v := range subroutineDec {
		if strings.Contains(line, v) {
			return true
		}
	}
	return false
}

func ParseSubroutineDec(lines []string, index int) (res []string, idx int) {
	res = append(res, "  <subroutineDec>")

	for i := index; i < len(lines); i++ {
		if strings.Contains(lines[i], "(") {
			res = append(res, "	"+lines[i])
			// parse parameter list
			l, resIdx := parseParameterList(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}
		if strings.Contains(lines[i], "{") {
			// subroutine body
			l, resIdx := parseSubroutineBody(lines, i)
			res = append(res, l...)
			idx = resIdx
			break
		}
		if strings.Contains(lines[i], "}") {
			idx = i
			break
		}
		res = append(res, "	"+lines[i])
	}

	res = append(res, "  </subroutineDec>")
	return
}

func parseParameterList(lines []string, index int) (res []string, idx int) {
	res = append(res, "	  <parameterList>")
	for i := index; i < len(lines); i++ {
		if strings.Contains(lines[i], ")") {
			res = append(res, "	  </parameterList>")
			res = append(res, "	"+lines[i])
			idx = i
			break
		}
		res = append(res, "		"+lines[i])
	}
	return
}

func parseSubroutineBody(lines []string, index int) (res []string, idx int) {
	res = append(res, "	<subroutineBody>")
	res = append(res, "	  "+lines[index])
	for i := index + 1; i < len(lines); i++ {
		if strings.Contains(lines[i], "}") {
			res = append(res, "	  "+lines[i])
			res = append(res, "	</subroutineBody>")
			idx = i
			break
		}
		if strings.Contains(lines[i], "var") {
			l, resIdx := parseVarDec(lines, i)
			res = append(res, l...)
			i = resIdx
			continue
		}
		if isStatement(lines[i]) {
			l, resIdx := parseStatements(lines, i)
			res = append(res, l...)
			i = resIdx
			continue
		}
		res = append(res, "		"+lines[i])
	}
	return
}

func parseVarDec(lines []string, index int) (res []string, idx int) {
	res = append(res, "	  <varDec>")
	for i := index; i < len(lines); i++ {
		if strings.Contains(lines[i], ";") {
			res = append(res, "		"+lines[i])
			res = append(res, "	  </varDec>")
			idx = i
			break
		}
		res = append(res, "		"+lines[i])
	}
	return
}

func containsOp(line string) bool {
	for _, v := range op {
		if strings.Contains(line, " "+v+" ") {
			return true
		}
	}
	return false
}

func containsUnaryOp(line string) bool {
	for _, v := range unaryOp {
		if strings.Contains(line, v) {
			return true
		}
	}
	return false
}

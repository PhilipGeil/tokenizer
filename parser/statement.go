package parser

import "strings"

func isStatement(line string) bool {
	for _, v := range statement {
		if strings.Contains(line, v) {
			return true
		}
	}
	return false
}

func parseStatements(lines []string, index int) (res []string, idx int) {
	res = append(res, "	  <statements>")
	for i := index; i < len(lines); i++ {
		if strings.Contains(lines[i], "}") {
			res = append(res, "	  </statements>")
			idx = i - 1
			break
		}
		// parse let statement
		if strings.Contains(lines[i], "let") {
			l, resIdx := parseLetStatement(lines, i)
			res = append(res, l...)
			i = resIdx
			continue
		}

		// parse if statement
		if strings.Contains(lines[i], " if ") {
			l, resIdx := parseIfStatement(lines, i)
			res = append(res, l...)
			i = resIdx + 1
			continue
		}

		// parse while statement
		if strings.Contains(lines[i], " while ") {
			l, resIdx := parseWhileStatement(lines, i)
			res = append(res, l...)
			i = resIdx + 1
			continue
		}

		// parse do statement
		if strings.Contains(lines[i], " do ") {
			l, resIdx := parseDoStatement(lines, i)
			res = append(res, l...)
			i = resIdx
			continue
		}

		// parse return statement
		if strings.Contains(lines[i], " return ") {
			l, resIdx := parseReturnStatement(lines, i)
			res = append(res, l...)
			i = resIdx
			continue
		}
	}
	return
}

func parseLetStatement(lines []string, index int) (res []string, idx int) {
	res = append(res, "		<letStatement>")
	for i := index; i < len(lines); i++ {

		// if semicolon is found, then it is the end of the let statement
		if strings.Contains(lines[i], ";") {
			res = append(res, "		  "+lines[i])
			res = append(res, "		</letStatement>")
			idx = i
			break
		}

		// if equals sign is found, then it is the end of the variable name
		if strings.Contains(lines[i], "=") {
			res = append(res, "		  "+lines[i])
			l, resIdx := parseExpression(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}

		// if square brackets are found here it means that it is an array
		if strings.Contains(lines[i], "[") {
			res = append(res, "		  "+lines[i])
			l, resIdx := parseExpression(lines, i+1)
			res = append(res, l...)
			i = resIdx + 1
			continue
		}
		res = append(res, "		  "+lines[i])
	}
	return
}

func parseIfStatement(lines []string, index int) (res []string, idx int) {
	res = append(res, "		<ifStatement>")
	for i := index; i < len(lines); i++ {

		// if curly brackets are found, then it is the end of the if statement unless it is an else statement
		if strings.Contains(lines[i], "}") && !strings.Contains(lines[i], "else") && !strings.Contains(lines[i+1], "else") {
			res = append(res, "		  "+lines[i])
			res = append(res, "		</ifStatement>")
			idx = i - 1
			break
		}

		// if parantheses are fount this is the beginning of the expression
		if strings.Contains(lines[i], "(") {
			res = append(res, "		  "+lines[i])
			l, resIdx := parseExpression(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}

		// if curly brackets are found, this is the beginning of the body
		if strings.Contains(lines[i], "{") {
			res = append(res, "		  "+lines[i])
			l, resIdx := parseStatements(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}

		// add whatever is left
		res = append(res, "		  "+lines[i])
	}
	return
}

func parseDoStatement(lines []string, index int) (res []string, idx int) {
	res = append(res, "		<doStatement>")
	for i := index; i < len(lines); i++ {

		// if semicolon is found, then it is the end of the do statement
		if strings.Contains(lines[i], ";") {
			res = append(res, "		  "+lines[i])
			res = append(res, "		</doStatement>")
			idx = i
			break
		}

		// if parantheses are found, then it is the beginning of the expression list
		if strings.Contains(lines[i], "(") {
			res = append(res, "		  "+lines[i])
			l, resIdx := parseExpressionList(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}
		res = append(res, "		  "+lines[i])
	}
	return
}

func parseWhileStatement(lines []string, index int) (res []string, idx int) {
	res = append(res, "		<whileStatement>")
	for i := index; i < len(lines); i++ {
		// if curly brackets are found, then it is the end of the while statement
		if strings.Contains(lines[i], "}") {
			res = append(res, "		  "+lines[i])
			res = append(res, "		</whileStatement>")
			idx = i - 1
			break
		}

		// if parantheses are found, then it is the beginning of the expression
		if strings.Contains(lines[i], "(") {
			res = append(res, "		  "+lines[i])
			l, resIdx := parseExpression(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}

		// if curly brackets are found, then it is the beginning of the body
		if strings.Contains(lines[i], "{") {
			res = append(res, "		  "+lines[i])
			l, resIdx := parseStatements(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}

		// add whatever is left
		res = append(res, "		  "+lines[i])
	}
	return
}

func parseReturnStatement(lines []string, index int) (res []string, idx int) {
	res = append(res, "		<returnStatement>")
	for i := index; i < len(lines); i++ {
		// if semicolon is found, then it is the end of the return statement
		if strings.Contains(lines[i], ";") {
			res = append(res, "		  "+lines[i])
			res = append(res, "		</returnStatement>")
			idx = i
			break
		}

		// if it contains a return keyword, then just add the line
		if strings.Contains(lines[i], " return ") {
			res = append(res, "		"+lines[i])
			continue
		}

		// this will be the value being returned if any
		res = append(res, "		  <expression>")
		res = append(res, "		  <term>")
		res = append(res, "		  "+lines[i])
		res = append(res, "		  </term>")
		res = append(res, "		  </expression>")
	}
	return
}

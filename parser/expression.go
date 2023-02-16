package parser

import "strings"

func parseExpression(lines []string, index int) (res []string, idx int) {
	res = append(res, "		<expression>")
	res = append(res, "		<term>")
	for i := index; i < len(lines); i++ {
		if strings.Contains(lines[i], " ; ") {
			res = append(res, "		</term>")
			res = append(res, "		</expression>")
			idx = i - 1
			break
		}
		if strings.Contains(lines[i], "]") {
			res = append(res, "		</term>")
			res = append(res, "		</expression>")
			res = append(res, "		"+lines[i])
			idx = i - 1
			break
		}
		if strings.Contains(lines[i], ")") {
			res = append(res, "		</term>")
			res = append(res, "		</expression>")
			res = append(res, "		"+lines[i])
			idx = i
			break
		}

		if containsUnaryOp(lines[i]) && strings.Contains(lines[i+1], "(") {
			res = append(res, "		"+lines[i])
			res = append(res, "		<term>")
			res = append(res, "		"+lines[i+1])
			l, resIdx := parseExpression(lines, i+2)
			res = append(res, l...)
			res = append(res, "		</term>")
			i = resIdx
			continue
		}
		if containsUnaryOp(lines[i]) && !strings.Contains(lines[i-1], "identifier") {
			res = append(res, "		"+lines[i])
			res = append(res, "		<term>")
			res = append(res, "		"+lines[i+1])
			res = append(res, "		</term>")
			i = i + 1
			continue
		}

		if containsOp(lines[i]) {
			res = append(res, "		</term>")
			res = append(res, "		"+lines[i])
			res = append(res, "		<term>")
			continue
		}
		if strings.Contains(lines[i], "(") && containsOp(lines[i-1]) {
			// parse expression
			res = append(res, "		"+lines[i])
			l, resIdx := parseExpression(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}
		if strings.Contains(lines[i], "[") {
			res = append(res, "		"+lines[i])
			l, resIdx := parseExpression(lines, i+1)
			res = append(res, l...)
			i = resIdx + 1
			continue
		}
		if strings.Contains(lines[i], "(") && strings.Contains(lines[i+1], "(") {
			res = append(res, "		"+lines[i])
			l, resIdx := parseExpression(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}
		if strings.Contains(lines[i], "(") && strings.Contains(lines[i-1], "(") {
			res = append(res, "		"+lines[i])
			l, resIdx := parseExpression(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}
		if strings.Contains(lines[i], "(") {
			res = append(res, "		"+lines[i])
			l, resIdx := parseExpressionList(lines, i+1)
			res = append(res, l...)
			i = resIdx
			continue
		}
		res = append(res, "		"+lines[i])

	}
	return
}

func parseExpressionList(lines []string, index int) (res []string, idx int) {
	res = append(res, "		<expressionList>")
	for i := index; i < len(lines); i++ {
		if strings.Contains(lines[i], ")") && strings.Contains(lines[i+1], ";") {
			res = append(res, "		</expressionList>")
			res = append(res, "		"+lines[i])
			idx = i
			break
		}
		if strings.Contains(lines[i], ",") {
			res = append(res, "		"+lines[i])
			continue
		}
		if strings.Contains(lines[i], "(") {
			res = append(res, "		<expression>")
			res = append(res, "		<term>")
			res = append(res, "		"+lines[i])
			continue
		}
		if strings.Contains(lines[i], ")") {
			// res = append(res, "		<expression>")
			// res = append(res, "		<term>")
			res = append(res, "		"+lines[i])
			res = append(res, "		</term>")
			continue
		}

		if containsUnaryOp(lines[i]) {
			res = append(res, "		"+lines[i])
			continue
		}

		if !containsUnaryOp(lines[i-1]) {
			res = append(res, "		<expression>")
		}
		res = append(res, "		<term>")
		res = append(res, "		"+lines[i])
		res = append(res, "		</term>")

		// check if next line is an op
		if containsOp(lines[i+1]) {
			res = append(res, "		"+lines[i+1])
			res = append(res, "		<term>")
			res = append(res, "		"+lines[i+2])
			res = append(res, "		</term>")
			res = append(res, "		</expression>")
			i = i + 2
			continue
		}
		res = append(res, "		</expression>")

	}
	return
}

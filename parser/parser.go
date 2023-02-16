package parser

import (
	"bufio"
	"os"
	"strings"
)

type Parser struct {
	Lines []string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseToken(input, output string) error {
	// Read all lines from file
	content, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Convert lines to string
	lines := strings.Split(string(content), "\r\n")

	// remove multiline comments
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "/*") {
			for j := i; j < len(lines); j++ {
				if strings.Contains(lines[j], "*/") {
					lines[i] = strings.Split(lines[j], "*/")[1]
					lines[j] = ""
					break
				}
				lines[j] = ""
			}
		}
	}
	linesToWrite := []string{"<tokens>"}

	for _, line := range lines {
		// remove comments
		line = strings.Split(line, "//")[0]

		// skip empty line
		if line == "" {
			continue
		}
		linesToWrite = append(linesToWrite, p.parseTokenLine(line)...)
	}

	linesToWrite = append(linesToWrite, "</tokens>")
	err = p.ParseClass(linesToWrite, output+".xml")
	if err != nil {
		panic(err)
	}
	return writeToXMLFile(linesToWrite, output+"T.xml")
}

func (p *Parser) ParseClass(lines []string, output string) error {
	linesToWrite := []string{"<class>"}
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "tokens") {
			continue
		}
		if strings.Contains(lines[i], "keyword") {
			if isClassVarDec(lines[i]) {
				res, idx := ParseClassVarDec(lines, i)
				linesToWrite = append(linesToWrite, res...)
				i = idx
				continue
			} else if isSubroutineDec(lines[i]) {
				res, idx := ParseSubroutineDec(lines, i)
				linesToWrite = append(linesToWrite, res...)
				i = idx
				continue
			}
		}
		linesToWrite = append(linesToWrite, lines[i])
	}
	linesToWrite = append(linesToWrite, "</class>")
	return writeToXMLFile(linesToWrite, output)

}

func writeToXMLFile(lines []string, name string) error {
	// Open file using os.Create()
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Write each line to the file
	for _, line := range lines {
		_, err := writer.WriteString(line + "\r\n")
		if err != nil {
			panic(err)
		}
	}

	// Flush the buffer to ensure everything is written to the file
	return writer.Flush()
}

func (p *Parser) parseTokenLine(line string) (lines []string) {
	// loop untill line is empty
	for line != "" {
		// check if line starts with a symbol
		if containsSymbol(string(line[0])) {
			lines = append(lines, p.parseSymbol(string(line[0])))
			line = line[1:]
			continue
		}

		// check if line starts with a keyword
		keyword := startsWithKeyword(line)
		if keyword != "" {
			lines = append(lines, p.parseKeyword(keyword))
			line = line[len(keyword):]
			continue
		}

		// check if line starts with a string
		if string(line[0]) == "\"" {
			// get string
			str := strings.Split(line, "\"")[1]
			lines = append(lines, p.parseString(str))
			line = line[len(str)+2:]
			continue
		}

		// check if line starts with a digit
		if isDigit(line[0]) {
			// get integer
			integer := getInteger(line)
			lines = append(lines, p.parseInteger(integer))
			line = line[len(integer):]
			continue
		}

		// check if line starts with an identifier
		identifier := getIdentifier(line)
		if identifier != "" {
			lines = append(lines, p.parseIdentifier(identifier))
			line = line[len(identifier):]
			continue
		}

		// if nothing matches, remove first char
		line = line[1:]
	}

	return
}

func getInteger(line string) string {
	var integer []rune
	for _, char := range line {
		if !isDigit(byte(char)) {
			return string(integer)
		}
		integer = append(integer, char)
	}
	return string(integer)
}

func getIdentifier(line string) string {
	var identifier []rune
	for _, char := range line {
		if containsSymbol(string(char)) {
			return string(identifier)
		}
		// return if char is space
		if string(char) == " " {
			return string(identifier)
		}
		// return if char is tab
		if string(char) == "\t" {
			return string(identifier)
		}
		identifier = append(identifier, char)
	}
	return string(identifier)
}

func isDigit(b byte) bool {
	return string(b) >= "0" && string(b) <= "9"
}

func (p *Parser) parseKeyword(line string) string {
	return stringToXML("keyword", line)
}

func (p *Parser) parseSymbol(line string) string {
	if line == "<" {
		line = "&lt;"
	} else if line == ">" {
		line = "&gt;"
	} else if line == "&" {
		line = "&amp;"
	}
	return stringToXML("symbol", line)
}

func (p *Parser) parseIdentifier(line string) string {
	return stringToXML("identifier", line)
}

func (p *Parser) parseInteger(line string) string {
	return stringToXML("integerConstant", line)
}

func (p *Parser) parseString(line string) string {
	return stringToXML("stringConstant", line)
}

func stringToXML(key, value string) string {
	return "<" + key + "> " + value + " </" + key + ">"
}

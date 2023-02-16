package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PhilipGeil/tokenizer/parser"
)

func main() {
	parser := parser.NewParser()

	files, err := loadFiles("ExpressionLessSquare")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println("Parsing file: " + file)

		ext := filepath.Ext(file)
		// remove file extension
		fileName := strings.Split(file, ext)[0]
		err = parser.ParseToken("ExpressionLessSquare/"+file, "ExpressionLessSquare/M_"+fileName)
		if err != nil {
			panic(err)
		}
	}

}

func loadFiles(path string) ([]string, error) {
	var lines []string

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".jack" {
			lines = append(lines, file.Name())
		}
	}

	return lines, nil
}

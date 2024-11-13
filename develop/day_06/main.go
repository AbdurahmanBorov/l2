package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	defaultDelimiter := "\t"

	if *delimiter != "" {
		defaultDelimiter = string((*delimiter)[0])
	}

	fieldIndices := make(map[int]bool)
	if *fields != "" {
		fieldsList := strings.Split(*fields, ",")
		for _, field := range fieldsList {
			fieldIndex := atoi(field)
			fieldIndices[fieldIndex-1] = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, defaultDelimiter) {
			if *separated {
				continue
			}

			fmt.Println(line)
			continue
		}

		fields := strings.Split(line, defaultDelimiter)

		var resultFields []string
		for index, field := range fields {
			if len(fieldIndices) == 0 || fieldIndices[index] {
				resultFields = append(resultFields, field)
			}
		}

		fmt.Println(strings.Join(resultFields, defaultDelimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка при чтении стандартного ввода:", err)
		os.Exit(1)
	}
}

func atoi(s string) int {
	result := 0
	for _, c := range s {
		result = result*10 + int(c-'0')
	}
	return result
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var hFlag bool
var rFlag bool
var kFlag int
var uFlag bool

func main() {
	flag.BoolVar(&hFlag, "h", false, "sort by human readable number (with suffixes)")
	flag.BoolVar(&rFlag, "r", false, "reverse sort")
	flag.IntVar(&kFlag, "k", 0, "sort by column (1-indexed)")
	flag.BoolVar(&uFlag, "u", false, "unique values")
	flag.Parse()

	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("No input file provided")
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(lines) == 0 {
		fmt.Println("No lines read from the file")
		return
	}

	if uFlag {
		lines = removeDuplicates(lines)
	}

	if hFlag {
		sort.SliceStable(lines, func(i, j int) bool {
			return convertToBytes(lines[i]) < convertToBytes(lines[j])
		})
	} else {
		if kFlag == 0 {
			sort.Strings(lines)
		} else {
			// Сортируем по указанной колонке
			sort.SliceStable(lines, func(i, j int) bool {
				colsI := strings.Fields(lines[i])
				colsJ := strings.Fields(lines[j])
				if kFlag-1 < len(colsI) && kFlag-1 < len(colsJ) {
					return colsI[kFlag-1] < colsJ[kFlag-1]
				}
				return false
			})
		}
	}

	if rFlag {
		reverse(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func convertToBytes(s string) int64 {
	re := regexp.MustCompile(`(\d+)([KMGTP]?)`)
	match := re.FindStringSubmatch(s)

	if len(match) < 3 {
		return 0
	}

	number, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		return 0
	}

	switch match[2] {
	case "K":
		return number * 1024
	case "M":
		return number * 1024 * 1024
	case "G":
		return number * 1024 * 1024 * 1024
	case "T":
		return number * 1024 * 1024 * 1024 * 1024
	default:
		return number
	}
}

func removeDuplicates(lines []string) []string {
	unique := make(map[string]struct{})
	var result []string

	for _, line := range lines {
		if _, exists := unique[line]; !exists {
			unique[line] = struct{}{}
			result = append(result, line)
		}
	}

	return result
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

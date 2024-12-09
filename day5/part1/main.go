package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	first, second string
}

func main() {
	f, err := os.Open("../input")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	rules := map[string][]string{}

	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		digits := strings.Split(line, "|")

		rules[digits[0]] = append(rules[digits[0]], digits[1])
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	lines := []string{}

	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}

	sum := 0

	for _, l := range lines {
		if right, middlePage := alreadyRight(l, rules); right {
			num, err := strconv.Atoi(middlePage)
			if err != nil {
				panic(err)
			}
			sum += num
		}

	}
	fmt.Println(sum)

}

func alreadyRight(updateline string, rules map[string][]string) (bool, string) {
	pages := strings.Split(updateline, ",")

	seen := []string{}

	for _, p := range pages {
		seen = append(seen, p)
		for _, rulePage := range rules[p] {
			if slices.Contains(seen, rulePage) {
				return false, ""
			}
		}
	}

	return true, pages[len(pages)/2]
}

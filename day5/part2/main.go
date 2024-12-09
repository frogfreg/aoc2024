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
	if err := sc.Err(); err != nil {
		panic(err)
	}

	unordered := []string{}

	for _, l := range lines {
		if !isOrdered(l, rules) {
			unordered = append(unordered, l)
		}
	}

	sum := 0
	for _, un := range unordered {

		num, err := strconv.Atoi(orderAndGetMiddle(un, rules))
		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Println(sum)

}

func orderAndGetMiddle(updateline string, rules map[string][]string) string {
	pages := strings.Split(updateline, ",")

	slices.SortFunc(pages, func(a, b string) int {
		if slices.Contains(rules[a], b) {
			return -1
		}
		if slices.Contains(rules[b], a) {
			return +1
		}
		return 0
	})

	return pages[len(pages)/2]

}

func isOrdered(updateline string, rules map[string][]string) bool {
	pages := strings.Split(updateline, ",")

	seen := []string{}

	for _, p := range pages {
		seen = append(seen, p)
		for _, rulePage := range rules[p] {
			if slices.Contains(seen, rulePage) {
				return false
			}
		}
	}

	return true
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	phrase := readString()
	abbr := abbreviation(phrase)
	fmt.Println(string(abbr))
}

func abbreviation(phrase string) (abbr []rune) {
	s := strings.Fields(phrase)
	for _, n := range s {
		r, _ := utf8.DecodeRuneInString(n)
		if unicode.IsLetter(r) == true {
			abbr = append(abbr, (unicode.ToUpper(r)))
		}
	}
	return abbr
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}

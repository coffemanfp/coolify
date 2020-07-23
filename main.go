package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

const vowels = `[aeiou]`

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	vowelsRe := regexp.MustCompile(vowels)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := []byte(strings.ToLower(s.Text()))

		vowerlsIndex := vowelsRe.FindAllIndex(word, -1)
		if len(vowerlsIndex) == 0 {
			fmt.Println(string(word))
			continue
		}

		vI := vowerlsIndex[rand.Intn(len(vowerlsIndex))][0]

		switch randBool() {
		case duplicateVowel:
			word = append(word[:vI+1], word[vI:]...)
		case removeVowel:
			word = append(word[:vI], word[vI+1:]...)
		}

		fmt.Println(string(word))
	}
}

func randBool() bool {
	return rand.Intn(2) == 0
}

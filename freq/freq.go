package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freq := make(map[string]int)

	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, word := range words {
			freq[strings.ToLower(word)] += 1
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return freq, nil
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}
	maxN, maxW := 0, ""
	for word, freq := range freqs {
		if freq > maxN {
			maxN, maxW = freq, word
		}
	}

	return maxW, nil
}

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}

	return maxWord(freqs)
}

func main() {
	absPath, _ := filepath.Abs("./practical_go/freq")
	f, err := os.Open(path.Join(absPath, "text.txt"))
	if err != nil {
		log.Fatalf("can't open the file: %s", err)
	}
	defer f.Close()

	fmt.Println(mostCommon(f))
}

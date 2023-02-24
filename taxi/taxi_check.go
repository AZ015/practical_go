package main

import (
	"bufio"
	"compress/bzip2"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type result struct {
	fileName string
	err      error
	match    bool
}

func sigWorker(fileName, signature string, ch chan<- result) {
	r := result{fileName: fileName}
	sig, err := fileSig(fileName)
	if err != nil {
		r.err = err
	} else {
		r.match = sig == signature
	}
	ch <- r
}

func fileSig(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, bzip2.NewReader(file))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), err
}

func parseSigFile(r io.Reader) (map[string]string, error) {
	sigs := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return nil, fmt.Errorf("bad line: %q", scanner.Text())
		}
		sigs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}

func main() {
	absPath, _ := filepath.Abs("./practical_go")
	rootDir := path.Join(absPath, "./taxi/tmp")
	file, err := os.Open(path.Join(rootDir, "sha256sum.txt"))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	sigs, err := parseSigFile(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	start := time.Now()
	ok := true
	ch := make(chan result)
	for name, signature := range sigs {
		fileName := path.Join(rootDir, name) + ".bz2"
		go sigWorker(fileName, signature, ch)
	}

	for range sigs {
		r := <-ch
		if r.err != nil {
			fmt.Fprintf(os.Stderr, "error: %s - %s\n", r.fileName, err)
			ok = false
			continue
		}

		if !r.match {
			ok = false
			fmt.Printf("error: %s mismatch\n", r.fileName)
		}
	}

	duration := time.Since(start)
	fmt.Printf("processed %d filed in %v\n", len(sigs), duration)
	if !ok {
		os.Exit(1)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const delimiter = "============"

var (
	sjisEncoder = japanese.ShiftJIS.NewEncoder()
	ngmap       = map[string]int{}
)

func main() {
	echoResult(os.Stdin, os.Stdout)
}

func echoResult(r io.Reader, w io.Writer) {
	err := echoHighlighted(r, w)
	if err != nil {
		log.Fatal(err)
	}

	if len(ngmap) == 0 {
		return
	}

	_, err = fmt.Fprintln(w, delimiter)
	if err != nil {
		log.Fatal(err)
	}

	err = echoUnicodeFormat(w)
	if err != nil {
		log.Fatal(err)
	}
}

func echoHighlighted(r io.Reader, w io.Writer) error {
	var (
		bufr  = bufio.NewReader(r)
		ngidx = 0
	)

	for {
		r, _, err := bufr.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		word := string(r)
		_, _, err = transform.String(sjisEncoder, word)
		if err != nil {
			_, err = fmt.Fprintf(w, "\x1b[31m%s\x1b[0m", word)
			if err != nil {
				return err
			}
			if _, ok := ngmap[word]; !ok {
				ngmap[word] = ngidx
				ngidx++
			}
			continue
		}
		_, err = fmt.Fprint(w, word)
		if err != nil {
			return err
		}
	}

	return nil
}

func echoUnicodeFormat(w io.Writer) error {
	var ngList = make([]string, len(ngmap))

	for ngWord, i := range ngmap {
		ngList[i] = ngWord
	}

	for _, ngWord := range ngList {
		_, err := fmt.Fprintf(w, "%#U\n", []rune(ngWord))
		if err != nil {
			return err
		}
	}

	return nil
}
